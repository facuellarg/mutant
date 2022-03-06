package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func MarshalToDynamoInput(v interface{}, tableName string) (*dynamodb.PutItemInput, error) {
	item, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		return nil, err
	}
	return &dynamodb.PutItemInput{
		Item:      item,
		TableName: &tableName,
	}, nil
}
