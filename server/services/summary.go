package services

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
)

// Convert ms to hour and minute format
func ConvertMsToHour(ms float64) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

func FormatDate(stringDate string) (time.Time, time.Time, error) {
	// Parse date
	sod, err := time.Parse("2006-01-02", stringDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	eod := sod.Add(24 * time.Hour)

	return sod, eod, nil
}

func PickWinner() uuid.UUID {
	var u models.User

	// Get user from db randomly
	database.DB.Order("RANDOM()").First(&u)

	return u.ID
}
