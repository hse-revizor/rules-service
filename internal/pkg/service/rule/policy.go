package rule

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
)

func (s *Service) CreatePolicy(ctx context.Context, policy *models.Policy) (*models.Policy, error) {
	return s.storage.CreatePolicy(ctx, policy)
}

func (s *Service) GetPolicyById(ctx context.Context, id uuid.UUID) (*models.Policy, error) {
	return s.storage.FindPolicyById(ctx, id)
}

func (s *Service) DeletePolicy(ctx context.Context, id uuid.UUID) (*models.Policy, error) {
	return s.storage.DeletePolicy(ctx, id)
}
