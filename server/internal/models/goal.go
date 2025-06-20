package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Goal struct {
	Timestamps
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Type       string    `json:"type"`
	Value      float64   `json:"value"`
	Name       string    `json:"name"`
	Unit       string    `json:"unit"`
	Comparison string    `json:"comparison"`
	Active     bool      `json:"active"`
}

func (g *Goal) BeforeCreate(_ *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}
