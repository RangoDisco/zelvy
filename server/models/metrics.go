package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Metrics struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Date         time.Time
	Steps        int
	KcalBurned   int
	KcalConsumed int
	Workouts     []Workout `gorm:"foreignKey:MetricsRefer"`
}

// Generates UUID before persist
func (m *Metrics) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
