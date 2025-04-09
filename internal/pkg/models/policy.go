package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Policy struct {
	ID        uuid.UUID    `gorm:"type:uuid;primary_key"`
	ProjectID string       `gorm:"column:project_id"`
	RulesIDs  StringStruct `gorm:"column:rules_ids;type:jsonb"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type StringStr struct {
	Value string `json:"value"`
}

type StringStruct []StringStr

func (t StringStruct) Value() (driver.Value, error) {
	if len(t) == 0 {
		return "[]", nil
	}
	return json.Marshal(t)
}

func (t *StringStruct) Scan(v interface{}) error {
	if v == nil {
		*t = StringStruct{}
		return nil
	}

	switch vt := v.(type) {
	case []byte:
		return json.Unmarshal(vt, &t)
	case string:
		return json.Unmarshal([]byte(vt), &t)
	default:
		return errors.New("invalid type for StringStruct scan")
	}
}

func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
