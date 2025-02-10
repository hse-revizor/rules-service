package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rule struct {
	Id uuid.UUID `gorm:"primaryKey;column:id"`

	TypeId string `gorm:"column:type_id"`
	Params Params `gorm:"column:params;type:jsonb"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// type RuleGroup struct {
// 	Id          uuid.UUID `gorm:"primaryKey;column:id"`
// 	Name        string    `gorm:"column:name"`
// 	Description string    `gorm:"column:description"`
// 	Rules       []Rule    `gorm:"many2many:rule_groups_rules"`

// 	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
// 	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
// }

type Params map[string]interface{}

func (p Params) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *Params) Scan(value interface{}) error {
	if value == nil {
		*p = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &p)
}

func (g *Rule) BeforeCreate(tx *gorm.DB) error {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}
