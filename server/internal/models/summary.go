package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Summary struct {
	Timestamps
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Date     time.Time `json:"date"`
	Metrics  []Metric  `gorm:"foreignKey:SummaryID"`
	Workouts []Workout `gorm:"foreignKey:SummaryID"`
	WinnerID uuid.UUID `gorm:"type:uuid; default:null"`
	Winner   User      `gorm:"foreignKey:WinnerID;references:ID; default:null"`
}

// Generates UUID before persist
func (s *Summary) BeforeCreate(_ *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
