package ingress

import (
	"djeurnie/api/internal/models"
)

type IngressService interface {
	all(tenant models.Tenant) (*models.IngressList, error)
	get(tenant models.Tenant, egressID string) (*models.Ingress, error)
}
