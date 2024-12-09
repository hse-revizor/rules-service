package rule

import (
	"context"
	"errors"
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"github.com/hse-revizor/rules-service/internal/pkg/storage/sql"
	"github.com/hse-revizor/rules-service/internal/utils/json"
)

// @throws: ErrRuleNotFound, ErrRuleExists
func (s *Service) CreateRule(ctx context.Context, rule *models.Rule) (*models.Rule, error) {
	created, err := s.storage.CreateRule(ctx, rule)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityExists):
			return nil, ErrRuleExists
		case errors.Is(err, sql.ErrForeignKey):
			return nil, ErrForeignKeyError
		default:
			return nil, err
		}
	}
	log.Debug(fmt.Sprintf("Created Rule: %s", json.ToColorJson(created)))
	return created, nil
}

type UpdateRuleInput struct {
	Id            uuid.UUID
}

// @throws: ErrRuleNotFound, ErrRuleExists
func (s *Service) UpdateRule(ctx context.Context, input *UpdateRuleInput) (*models.Rule, error) {
	rule, err := s.storage.UpdateRule(ctx, &models.Rule{
		Id:            input.Id,
	})
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("Updated Rule: %s", json.ToColorJson(rule)))
	return rule, nil
}

// @throws: ErrRuleNotFound
func (s *Service) DeleteRule(ctx context.Context, ruleId uuid.UUID) (*models.Rule, error) {
	model, err := s.storage.DeleteRule(ctx, ruleId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrRuleNotFound
		default:
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("Deleted Rule Account: %s", json.ToColorJson(model)))

	return model, nil
}

func (s *Service) GetRuleById(ctx context.Context, ruleId uuid.UUID) (*models.Rule, error) {
	rule, err := s.storage.FindRuleById(ctx, ruleId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrRuleNotFound
		default:
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("Found Rule: %s", json.ToColorJson(rule)))

	return rule, nil

}


type GetAllRulesOutput struct {
	Rules []*models.Rule
	Count int32
}

func (s *Service) GetAllRules(ctx context.Context, skip, limit int) (*GetAllRulesOutput, error) {
	rules, err := s.storage.GetAllRules(ctx, sql.GetAllRulesPayload{
		PaginationInput: &sql.PaginationInput{
			Limit: limit,
			Skip:  skip,
		},
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrRuleNotFound
		default:
			return nil, err
		}
	}
	log.Debug(fmt.Sprintf("Found Rule: %s", json.ToColorJson(rules)))

	return &GetAllRulesOutput{
		Rules: rules.Rules,
		Count: rules.Count,
	}, nil

}