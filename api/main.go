package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handleError(err error) (events.APIGatewayProxyResponse, error) {
	log.Fatalf("%s", err)
	return events.APIGatewayProxyResponse{}, err
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// parse request

	// get egress from dynamodb

	// marshal egress

	// return marshalled egress as an api gateway response
}

func main() {
	lambda.Start(handler)
}

/*

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
*/
