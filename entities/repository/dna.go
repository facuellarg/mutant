package repository

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type (
	DNA struct {
		Sequence []string `json:"sequence"`
		IsMutant bool     `json:"is_mutant"`
		ID       string   `json:"id"`
	}
	DnaRepositoryI interface {
		Save(DNA) error
		GetAll() ([]DNA, error)
		GetMutantsCount() (int, error)
		GetHumansCount() (int, error)
	}
	dnaRepository struct {
		awsConnection *dynamodb.DynamoDB
		TableName     string
	}
)

const (
	TABLE_NAME = "Mutant"
)

func NewDnaRepository(awsConnection *dynamodb.DynamoDB) DnaRepositoryI {
	r := &dnaRepository{awsConnection, "Mutant"}
	r.createMutantTable()
	return r
}

func (dr *dnaRepository) Save(dna DNA) error {
	item, err := dynamodbattribute.MarshalMap(dna)
	if err != nil {
		return err
	}
	_, err = dr.awsConnection.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &dr.TableName,
	})
	return err
}

func (dt *dnaRepository) GetAll() ([]DNA, error) {
	return nil, nil
}

func (dt *dnaRepository) GetMutantsCount() (int, error) {

	filt := expression.Name("is_mutant").Equal(expression.Value(true))
	return dt.getCountMutantFilter(filt)
}

func (dt *dnaRepository) GetHumansCount() (int, error) {
	filt := expression.Name("is_mutant").Equal(expression.Value(false))
	return dt.getCountMutantFilter(filt)
}

func (dt *dnaRepository) getCountMutantFilter(filt expression.ConditionBuilder) (int, error) {
	proj := expression.NamesList(expression.Name("is_mutant"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		return 0, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TABLE_NAME),
	}

	result, err := dt.awsConnection.Scan(params)
	if err != nil {
		return 0, err
	}
	return int(*result.Count), nil
}

func (dt *dnaRepository) createMutantTable() {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(TABLE_NAME),
	}

	_, err := dt.awsConnection.CreateTable(input)
	if err != nil && !strings.Contains(err.Error(), "Table already exists: Mutant") {
		log.Fatalf("Giot error calling CreateTable: %s", err)
	}
}
