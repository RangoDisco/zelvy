package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Metric struct {
	Timestamps
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	SummaryID uuid.UUID
	GoalID    uuid.UUID `gorm:"type:uuid; default:null"`
	Goal      Goal      `gorm:"foreignKey:GoalID;references:ID; default:null"`
	Success   bool      `gorm:"default:false"`
	Disabled  bool      `gorm:"default:false"`
}

func (m *Metric) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
