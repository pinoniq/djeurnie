package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type QueryByHashKeyCondition struct {
	TableName string
	KeyName   string
	KeyValue  string
}

func (db *DBAL) QueryAllByHashKey(qc *QueryByHashKeyCondition) (*dynamodb.QueryOutput, error) {
	out, err := db.Client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(qc.TableName),
		KeyConditionExpression: aws.String(qc.KeyName + " = :Key"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":Key": &types.AttributeValueMemberS{Value: qc.KeyValue},
		},
	})

	return out, err
}
