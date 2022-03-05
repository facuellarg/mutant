package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var region = "us-east-2"

func Dynamodb() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(
		session.NewSession(&aws.Config{
			Region: aws.String(region),
		}),
	))
}
