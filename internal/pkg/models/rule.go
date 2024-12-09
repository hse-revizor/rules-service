package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rule struct {
	Id uuid.UUID `gorm:"primaryKey;column:id"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (g *Rule) BeforeCreate(tx *gorm.DB) error {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}
