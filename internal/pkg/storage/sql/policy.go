package sql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hse-revizor/rules-service/internal/pkg/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreatePolicy(ctx context.Context, model *models.Policy) (*models.Policy, error) {
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

func (s *Storage) FindPolicyById(ctx context.Context, id uuid.UUID) (*models.Policy, error) {
	intent := new(models.Policy)
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.
		Model(&models.Policy{}).
		Where("id = ?", id).
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

func (s *Storage) DeletePolicy(ctx context.Context, id uuid.UUID) (*models.Policy, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)

	deletedIntent := new(models.Policy)
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
