package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
	Tenant string `json:"tenant"`
}

func HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawTenant := c.Locals("tenant")

		if rawTenant == nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Tenant not found")
		}

		tenant := rawTenant.(models.Tenant)

		return c.JSON(HealthCheckResponse{
			Status: "ok",
			Tenant: tenant.Id,
		})
	}
}
