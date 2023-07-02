package main

import (
	"context"
	"djeurnie/api/cmd/functions/core/handlers"
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberAdapter.FiberLambda

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Djeurnie Core",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	app.Use(middleware.TenantMiddleware())

	app.Get("/", handlers.HealthCheck())

	if helpers.IsLambda() {
		fiberLambda = fiberAdapter.New(app)
		lambda.Start(Handler)
	} else {
		app.Listen(":3001")
	}

}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return fiberLambda.ProxyWithContextV2(ctx, request)
}
