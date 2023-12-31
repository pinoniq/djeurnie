package middleware

import (
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/models"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gofiber/fiber/v2"
)

// Hack our way around the aws package not passing along the context to the fiber http request
func getAPIGatewayContextV2(c *fiber.Ctx) (events.APIGatewayV2HTTPRequestContext, error) {
	reqHeaders := c.GetReqHeaders()

	apiGwContextHeader, ok := reqHeaders["X-Golambdaproxy-Apigw-Context"]

	if !ok {
		return events.APIGatewayV2HTTPRequestContext{}, errors.New("No context header in request")
	}

	context := events.APIGatewayV2HTTPRequestContext{}
	err := json.Unmarshal([]byte(apiGwContextHeader), &context)
	if err != nil {
		return events.APIGatewayV2HTTPRequestContext{}, err
	}

	return context, nil
}

func TenantMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		// default tenantId, useful for local use and testing
		// TODO: remove this default value when in production
		tenantId := "5cdaedd3-0cc7-4d9c-899c-6dc4e6b717ac"

		if helpers.IsLambda() {
			apiGwContext, err := getAPIGatewayContextV2(c)

			if err != nil {
				return err
			}

			tenantId = apiGwContext.DomainPrefix
		}

		c.Locals("tenant", models.Tenant{
			Id: tenantId,
		})

		return c.Next()
	}
}
