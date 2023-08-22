package main

import (
	"context"
	"djeurnie/api/cmd/functions/core/handlers"
	"djeurnie/api/internal/database"
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/middleware"
	apiRoutes "djeurnie/api/internal/service/api_routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberAdapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"os"
)

var fiberLambda *fiberAdapter.FiberLambda

var inLambda = helpers.IsLambda()

func main() {
	if !inLambda {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	resourceGroupId := os.Getenv("RESOURCE_GROUP_ID")

	svc := database.GetPlanetScalSession()
	apiRoutesService := apiRoutes.NewPlanetScaleDbService(svc)

	factory := handlers.NewFactory(svc)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Use(
		recover.New(),
		logger.New(logger.Config{
			// For more options, see the Config section
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}),
	)
	app.Use(middleware.DeciderOfEncodings())

	app.Get("/healthcheck", handlers.HealthCheck)

	// routes
	routes := apiRoutesService.AllForResourceGroup(resourceGroupId)
	for _, route := range routes.Items {
		app.Add(route.Method, route.Path, factory.CreateHandler(&route))
	}

	if inLambda {
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
