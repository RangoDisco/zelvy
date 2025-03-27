package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Username    string    `gorm:"type:varchar(75)" json:"username"`
	DiscordID   string    `gorm:"type:varchar(75);unique_index" json:"discordID"`
	PaypalEmail string    `gorm:"type:varchar(75);unique_index" json:"paypalEmail"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Generate UUID before persist
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
