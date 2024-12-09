package rule

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"github.com/hse-revizor/rules-service/internal/pkg/storage/sql"
)

type storage interface {
	CreateRule(context.Context, *models.Rule) (*models.Rule, error)
	FindRuleById(context.Context, uuid.UUID) (*models.Rule, error)
	UpdateRule(context.Context, *models.Rule) (*models.Rule, error)
	DeleteRule(context.Context, uuid.UUID) (*models.Rule, error)
	GetAllRules(ctx context.Context, input sql.GetAllRulesPayload) (*sql.GetAllRulesOutput, error)
}
type Service struct {
	storage      storage
}

func New(storage storage) *Service {
	return &Service{storage: storage}
}
