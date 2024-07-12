package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workout struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string
	KcalBurned   int
	Type         string
	Duration     int
	MetricsRefer uuid.UUID
}

// Generates UUID before persist
func (w *Workout) BeforeCreate(tx *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}
