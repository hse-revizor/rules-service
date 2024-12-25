package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rule struct {
	Id          uuid.UUID  `gorm:"primaryKey;column:id"`
	FilePath    string     `gorm:"column:file_path"`
	Item        string     `gorm:"column:item"`
	ShouldBe    string     `gorm:"column:should_be"`
	Type        RuleType   `gorm:"column:type"`
	WorkspaceId string     `gorm:"column:workspace_id"`
	
	CreatedAt   *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (g *Rule) BeforeCreate(tx *gorm.DB) error {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}
