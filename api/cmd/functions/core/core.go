package main

import (
	"context"
	"djeurnie/api/cmd/functions/core/handlers"
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/middleware"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var fiberLambda *fiberAdapter.FiberLambda

func encode[T any](c *fiber.Ctx, res T) error {
	accepts := c.Accepts("application/json", "application/x-protobuf", "application/xml")

	if accepts == "" {
		return errors.New("missing accept header")
	}

	switch accepts {
	case "application/json":
		return c.JSON(res)

	case "application/xml":
		return c.XML(res)
	}

	return errors.New("unsupported accept header")
}

func wrap[T any](handler func(c *fiber.Ctx) (T, error)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		res, err := handler(c)
		if err != nil {
			return c.SendStatus(fiber.StatusNotAcceptable)
		}

		return encode(c, res)
	}
}

func main() {
	app := fiber.New()

	app.Use(
		recover.New(),
		logger.New(logger.Config{
			// For more options, see the Config section
			Format: "[${time}] ${status} - ${latency} ${method} ${path}",
		}),
	)
	app.Use(middleware.TenantMiddleware())

	app.Get("/healthcheck", wrap(handlers.HealthCheck))

	if helpers.IsLambda() {
		fiberLambda = fiberAdapter.New(app)
		lambda.Start(Handler)
	} else {
		err := app.Listen(":3001")
		if err != nil {
			return
		}
	}

}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyV2(request)
}
