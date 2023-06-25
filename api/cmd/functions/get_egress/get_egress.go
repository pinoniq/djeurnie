package main

import (
	"context"
	"djeurnie/api/internal/database"
	transport "djeurnie/api/internal/transport/lambda"
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
	user, err := service.GetUserByID(ctx, userID)
	if err != nil {
		return transport.SendError(500, err)
	} else user == nil {
		return transport.SendError(404, "user not found")
	}

	return transport.Send(200, user)
}

func main() {
	lambda.Start(handler)
}