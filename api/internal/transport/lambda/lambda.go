package transport

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type Response = events.APIGatewayV2HTTPResponse

type ErrorMessage struct {
	Message string `json:"message"`
}

func Send(code int, body interface{}) (Response, error) {
	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return Response{
			StatusCode: code,
		}, err
	}

	return Response{
		StatusCode: code,
		Body:       string(bodyBytes),
	}, nil
}

func SendValidationError(code int, message string) (Response, error) {
	errorMessage := ErrorMessage{
		Message: message,
	}

	body, err := json.Marshal(errorMessage)

	if err != nil {
		return Response{
			StatusCode: code,
		}, err
	}

	return Response{
		StatusCode: code,
		Body:       string(body),
	}, nil
}

func SendError(code int, err error) (Response, error) {
	return SendValidationError(code, err.Error())
}
