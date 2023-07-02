package main

import (
	"context"
	"djeurnie/api/internal/database"
	transport "djeurnie/api/internal/transport/lambda"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request = events.APIGatewayV2HTTPRequest
type Response = events.APIGatewayV2HTTPResponse

var db = database.GetDynamodbSession()

func handler(ctx context.Context, r Request) (Response, error) {
	egressId := r.PathParameters["egressId"]
	if egressId == "" {
		return transport.SendValidationError(400, "egressId is empty")
	}

	// all the actual logic happens in that method call
	return transport.SendError(404, errors.New("not implemented"))
}

func main() {
	lambda.Start(handler)
}
