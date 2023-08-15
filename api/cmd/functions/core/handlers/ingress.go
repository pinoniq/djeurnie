package handlers

import (
	"djeurnie/api/internal/models"
	"djeurnie/api/internal/service/ingress"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	List(c *fiber.Ctx, tenant models.Tenant) (*ingressListResponse, error)
	Get(c *fiber.Ctx, tenant models.Tenant) (*ingressResponse, error)
}

type handler struct {
	Ingress ingress.Service
}

type ingressResponse struct {
	Id          string `json:"id" xml:"id"`
	DisplayName string `json:"displayName" xml:"displayName"`
}

type ingressListResponse struct {
	Status  string            `json:"status" xml:"status"`
	Tenant  string            `json:"tenant" xml:"tenant"`
	Ingress []ingressResponse `json:"items" xml:"items"`
}

func (h *handler) List(c *fiber.Ctx, tenant models.Tenant) (*ingressListResponse, error) {
	listRes := ingressListResponse{
		Status: "ok",
		Tenant: tenant.Id,
	}

	tenantsIngress := h.Ingress.All(tenant)

	for _, item := range tenantsIngress.Items {
		listRes.Ingress = append(listRes.Ingress, ingressResponse{
			Id:          item.Id,
			DisplayName: item.DisplayName,
		})
	}

	return &listRes, nil
}

func (h *handler) Get(c *fiber.Ctx, tenant models.Tenant) (*ingressResponse, error) {
	ingressId := c.Params("Id")

	item, err := h.Ingress.Get(tenant, ingressId)

	if err != nil {
		return nil, err
	}

	return &ingressResponse{
		Id:          item.Id,
		DisplayName: item.DisplayName,
	}, nil
}

func NewIngressHandler(ingress *ingress.Service) Handler {
	return &handler{Ingress: *ingress}
}
