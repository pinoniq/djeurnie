package handlers

import (
	"djeurnie/api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type IngressResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type IngressListResponse struct {
	Status  string            `json:"status" xml:"status"`
	Tenant  string            `json:"tenant" xml:"tenant"`
	Ingress []IngressResponse `json:"items" xml:"items"`
}

func IngressList(c *fiber.Ctx, tenant models.Tenant) (*IngressListResponse, error) {
	listRes := IngressListResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}

	listRes.Ingress = append(listRes.Ingress, IngressResponse{
		Id:          "ingress-0",
		DisplayName: "Jeroen is cool",
	})
	listRes.Ingress = append(listRes.Ingress, IngressResponse{
		Id:          "ingress-1",
		DisplayName: "Katleen ook",
	})

	return &listRes, nil
}
