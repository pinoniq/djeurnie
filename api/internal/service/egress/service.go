package egress

import (
	"context"
	"djeurnie/api/internal/models"
)

type EgressService interface {
	GetEgressByID(ctx context.Context, egressID string) (*models.Egress, error)
}
