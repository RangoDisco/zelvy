package services

import (
	"server/config/database"
	"server/internal/models"
)

func FetchGoals() ([]*models.Goal, error) {
	var goals []*models.Goal
	if err := database.GetDB().Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}
