package services

import (
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
)

func FetchGoals() ([]*models.Goal, error) {
	var goals []*models.Goal
	if err := database.GetDB().Find(&goals).Error; err != nil {
		return nil, err
	}
	return goals, nil
}
