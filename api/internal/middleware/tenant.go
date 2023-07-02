package middleware

import (
	"djeurnie/api/internal/helpers"
	"djeurnie/api/internal/models"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gofiber/fiber/v2"
)

func TenantMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		// default tenantId, useful for local use and testing
		// TODO: remove this default value when in production
		tenantId := "00000000-0000-0000-0000-000000000000"

		if helpers.IsLambda() {
			apiGwContext, ok := core.GetAPIGatewayV2ContextFromContext(c.Context())

			if ok {
				tenantId = apiGwContext.DomainPrefix
			}
		}

		c.Locals("tenant", models.Tenant{
			Id: tenantId,
		})

		return c.Next()
	}
}
