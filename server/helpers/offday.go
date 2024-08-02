package helpers

import (
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"time"
)

func FetchByGoalAndDate(goalType string) (*models.Offday, error) {
	var offDay models.Offday
	today := time.Now().Format("2006-01-02")
	if err := database.DB.First(&offDay, "type = ? AND date = ?", goalType, today).Error; err != nil {
		return nil, err
	}

	return &offDay, nil
}
