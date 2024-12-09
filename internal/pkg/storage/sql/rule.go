package sql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginationPayload struct {
	Limit int
	Token int64
}

func (s *Storage) CreateRule(ctx context.Context, model *models.Rule) (*models.Rule, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(&model).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return nil, ErrEntityExists
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return nil, ErrForeignKey
		default:
			return nil, err
		}
	}

	return model, nil
}
func (s *Storage) FindRuleById(ctx context.Context, id uuid.UUID) (*models.Rule, error) {
	intent := new(models.Rule)
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.
		Model(&models.Rule{}).Where("id = ?", id).
		First(intent).
		Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrEntityNotFound
		default:
			return nil, err
		}
	}

	return intent, nil
}


func (s *Storage) UpdateRule(ctx context.Context, model *models.Rule) (*models.Rule, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Clauses(clause.Returning{}).Model(&model).Where("id = ?", model.Id).Updates(model)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return nil, ErrEntityNotFound
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return nil, ErrEntityExists
		default:
			return nil, result.Error
		}
	}
	

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return model, nil
}

func (s *Storage) DeleteRule(ctx context.Context, id uuid.UUID) (*models.Rule, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)

	deletedIntent := new(models.Rule)
	result := tr.Clauses(clause.Returning{}).Where("id = ?", id).Delete(deletedIntent)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrEntityNotFound
		}

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return deletedIntent, nil
}
func (s *Storage) GetRulesById(ctx context.Context, ids []uuid.UUID) ([]*models.Rule, error) {
	res := make([]*models.Rule, 0)
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Model(&models.Rule{}).Where("id  in ?", ids).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PaginationInput struct {
	Limit int
	Skip  int
}

type GetAllRulesPayload struct {
	*PaginationInput
}
type GetAllRulesOutput struct {
	Rules []*models.Rule
	Count int32
}

func (s *Storage) GetAllRules(ctx context.Context, input GetAllRulesPayload) (*GetAllRulesOutput, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var out GetAllRulesOutput
	var count int64
	err := tr.
		Model(&models.Rule{}).
		Count(&count).
		Limit(input.Limit).
		Offset(input.Skip).
		Find(&out.Rules).
		Error
	if err != nil {
		return nil, err
	}
	out.Count = int32(count)
	return &out, nil
}