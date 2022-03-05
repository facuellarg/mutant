package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type (
	DNA struct {
		Sequence []string
		IsMutant bool
		ID       int
	}
	DnaRepositoryI interface {
		Save(DNA) error
		GetAll() ([]DNA, error)
		GetMutants() ([]DNA, error)
		GetHumans() ([]DNA, error)
	}
	dnaRepository struct {
		awsConnection *dynamodb.DynamoDB
	}
)

func NewDnaRepository(awsConnection *dynamodb.DynamoDB) DnaRepositoryI {
	return &dnaRepository{awsConnection}
}

func (dr *dnaRepository) Save(dna DNA) error {
	item, err := dynamodbattribute.MarshalMap(dna)
	if err != nil {
		return err
	}
	_, err = dr.awsConnection.PutItem(&dynamodb.PutItemInput{
		Item: item,
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
