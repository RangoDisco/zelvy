package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Summary struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Date time.Time `json:"date"`
	//Steps           int       `json:"steps"`
	//KcalBurned      int       `json:"kcalBurned"`
	//KcalConsumed    int       `json:"kcalConsumed"`
	//MilliliterDrank int       `json:"milliliterDrank"`
	Metrics  []Metric  `gorm:"foreignKey:SummaryID"`
	Workouts []Workout `gorm:"foreignKey:SummaryID"`
}

// Generates UUID before persist
func (s *Summary) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
