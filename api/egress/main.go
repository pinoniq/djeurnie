package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"net/http"
	"net/url"
)

func getDynamodbSession() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg, func(options *dynamodb.Options) {
		options.EnableAcceptEncodingGzip = true
	})
}

func handleError(err error) (events.APIGatewayProxyResponse, error) {
	log.Fatalf("%s", err)
	return events.APIGatewayProxyResponse{}, err
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	svc := getDynamodbSession()
	rawTableName, foundModel := request.PathParameters["model"]
	rawId, foundId := request.PathParameters["id"]
	revision := "current"

	if !foundModel || !foundId {
		return handleError(errors.New("No model or id defined."))
	}

	tableName, err := url.QueryUnescape(rawTableName)
	if nil != err {
		return handleError(err)
	}

	id, err := url.QueryUnescape(rawId)
	if nil != err {
		return handleError(err)
	}

	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id":       &types.AttributeValueMemberS{Value: id},
			"revision": &types.AttributeValueMemberS{Value: revision},
		},
	})
	if err != nil {
		return handleError(err)
	}
	if len(out.Item) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
		}, nil
	}

	jsonStr, err := json.Marshal(out.Item)

	if err != nil {
		return handleError(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonStr),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
