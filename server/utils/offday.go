package utils

import (
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"time"
)

func FetchByGoalAndDate(goalId uuid.UUID) *models.Offday {
	var offDay models.Offday
	sod := time.Now().Truncate(24 * time.Hour)
	eod := sod.Add(24 * time.Hour)
	if err := database.DB.First(&offDay, "goal_id = ? AND day >= ? AND day < ?", goalId, sod, eod).Error; err != nil {
		return nil
	}

	return &offDay
}
