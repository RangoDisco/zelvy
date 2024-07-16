package models

import (
	"github.com/google/uuid"
	"time"
)

type Offday struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Day    time.Time `json:"day"`
	GoalID uuid.UUID `gorm:"type:uuid;not null"`
	Goal   Goal      `gorm:"foreignKey:GoalID;references:ID"`
}
