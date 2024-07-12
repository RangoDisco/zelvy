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
	Duration     int       `json:"duration"`
	MetricsRefer uuid.UUID
}

// Generates UUID before persist
func (w *Workout) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}
