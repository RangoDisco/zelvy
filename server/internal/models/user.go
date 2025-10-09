package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Timestamps
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Username    string    `gorm:"type:varchar(75)" json:"username"`
	DiscordID   string    `gorm:"type:varchar(75);unique_index" json:"discordID"`
	PaypalEmail string    `gorm:"type:varchar(75);unique_index" json:"paypalEmail"`
	Picture     string    `gorm:"type:varchar(255)" json:"picture"`
}

// Generate UUID before persist
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
