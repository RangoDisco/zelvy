package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workout struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string    `json:"name"`
	KcalBurned   int       `json:"kcalBurned"`
	ActivityType string    `json:"activityType"`
	Duration     float64   `json:"duration"`
	SummaryID    uuid.UUID
}

// Generates UUID before persist
func (w *Workout) BeforeCreate(_ *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}
