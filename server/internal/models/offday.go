package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Offday struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Day    time.Time `json:"day"`
	GoalID uuid.UUID `gorm:"type:uuid;not null"`
	Goal   Goal      `gorm:"foreignKey:GoalID;references:ID"`
}

func (o *Offday) BeforeCreate(_ *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
