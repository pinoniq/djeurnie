package main

import (
	"context"
	"djeurnie/api/cmd/functions/core/handlers"
	"djeurnie/api/internal/database"
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

func main() {
	app := fiber.New()

	app.Use(
		recover.New(),
		logger.New(logger.Config{
			// For more options, see the Config section
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}),
	)
	app.Use(middleware.DeciderOfEncodings())
	app.Use(middleware.TenantMiddleware())

	app.Get("/healthcheck", transport.WrapEncodingWithTenant(handlers.HealthCheck))

	svc := database.GetDynamodbSession()

	// Ingress
	ingressService := ingress.NewDynamoDbService(svc, "ingress")
	ingressHandler := handlers.NewIngressHandler(&ingressService)
	app.Get("/ingress", transport.WrapEncodingWithTenant(ingressHandler.List))

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
