package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Metric struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	SummaryID uuid.UUID
	Goal      *Goal `json:"goal"`
	GoalID    uuid.UUID
}

func (m *Metric) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
