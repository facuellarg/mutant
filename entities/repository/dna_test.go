package repository_test

import (
	"errors"
	"testing"
	"tests/entities/repository"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

func (m *mockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	return nil, args.Error(0)
}

func TestSaveWithError(t *testing.T) {

	assert := assert.New(t)
	mockCon := new(mockDynamoDBClient)

	dna := repository.DNA{}
	input, _ := repository.MarshalToDynamoInput(dna, "Mutant")
	errOutput := errors.New("put tiem error")
	mockCon.On("PutItem", input).Return(errOutput)

	r := repository.NewDnaRepository(mockCon)

	err := r.Save(dna)
	assert.ErrorIs(err, errOutput)

}
