package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
	Tenant string `json:"tenant"`
}

type HealthCheckErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(c *fiber.Ctx) error {
	rawTenant := c.Locals("tenant")

	if rawTenant == nil {
		return c.JSON(HealthCheckErrorResponse{
			Status:  "error",
			Message: "No tenant found",
		})
	}

	tenant := rawTenant.(models.Tenant)

	return c.JSON(HealthCheckResponse{
		Status: "ok",
		Tenant: tenant.Id,
	})
}
