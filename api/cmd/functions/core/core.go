package main

import (
	"context"
	"djeurnie/api/cmd/functions/core/handlers"
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/middleware"
	"djeurnie/api/internal/service/ingress"
	"djeurnie/api/internal/transport"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var fiberLambda *fiberAdapter.FiberLambda

type Env struct {
	ingress ingress.IngressService
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
	app.Use(middleware.DeciderOfEncodings())
	app.Use(middleware.TenantMiddleware())

	app.Get("/healthcheck", transport.WrapEncodingWithTenant(handlers.HealthCheck))

	app.Get("/ingress", transport.WrapEncodingWithTenant(handlers.IngressList))

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
