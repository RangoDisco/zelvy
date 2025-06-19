package services

import (
	"time"

	"server/config/database"
	"server/internal/models"

	"github.com/google/uuid"
)

func IsOff(goalId uuid.UUID) bool {
	var offDay models.Offday
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)
	if err := database.GetDB().First(&offDay, "goal_id = ? AND day >= ? AND day < ?", goalId, sod, eod).Error; err != nil {
		return false
	}

	if offDay.ID != uuid.Nil {
		return true
	}

	return false
}
