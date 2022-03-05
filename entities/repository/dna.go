package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
		GetMutants() ([]DNA, error)
		GetHumans() ([]DNA, error)
	}
	dnaRepository struct {
		awsConnection *dynamodb.DynamoDB
		TableName     string
	}
)

func NewDnaRepository(awsConnection *dynamodb.DynamoDB) DnaRepositoryI {
	return &dnaRepository{awsConnection, "Mutant"}
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

func (dt *dnaRepository) GetMutants() ([]DNA, error) {
	return nil, nil
}

func (dt *dnaRepository) GetHumans() ([]DNA, error) {
	return nil, nil
}
