package database

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DBAL struct {
	Client dynamodb.Client
}
