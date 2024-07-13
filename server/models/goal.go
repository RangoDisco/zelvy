package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Goal struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Type  string    `json:"type"`
	Value int       `json:"value"`
}

func (g *Goal) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = uuid.New()
	return
}
