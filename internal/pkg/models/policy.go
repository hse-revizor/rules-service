package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Policy struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id"`
	ProjectID string    `gorm:"column:project_id"`
	RulesIDs  []string  `gorm:"column:rules_ids;type:text[]"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if p.Id == uuid.Nil {
		p.Id = uuid.New()
	}
	return nil
}
