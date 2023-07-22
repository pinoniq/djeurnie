package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type IngressResponse struct {
	Status  string `json:"status"`
	Tenant  string `json:"tenant"`
	Message string `json:"message"`
}

func IngressList(c *fiber.Ctx, tenant models.Tenant) (*IngressResponse, error) {
	return &IngressResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}, nil
}
