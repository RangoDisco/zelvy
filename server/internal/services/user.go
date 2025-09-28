package services

import (
	"github.com/google/uuid"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/models"
	"time"
)

// UpsertUser tries to find a user by its discord's id, in case it exists, updates the mail, otherwise creates a new user
func UpsertUser(body *pb_usr.AddUserRequest) error {

	eU, err := findExistingUser(body)
	if err != nil {
		return err
	}

	if eU != nil {
		err = updateUserEmail(eU, body.PaypalEmail)
		if err != nil {
			return err
		}
		return nil
	}

	iM := convertToInputModel(body)

	if err = database.GetDB().Save(&iM).Error; err != nil {
		return err
	}

	return nil
}

func GetWinnersBetweenDates(sod, eod string, limit int64) ([]*pb_usr.WinnerViewModel, error) {
	var winners []*pb_usr.WinnerViewModel
	err := database.GetDB().Raw(""+
		"SELECT u.username, u.picture, count(u.id) as wins FROM summaries s INNER JOIN users u ON s.winner_id = u.id WHERE s.date >= ? AND s.date <= ? AND s.deleted_at IS null GROUP BY u.username, u.picture ORDER BY wins DESC LIMIT ?",
		sod, eod, limit).Scan(&winners).Error

	if err != nil {
		return nil, err
	}

	return winners, nil
}

func findExistingUser(body *pb_usr.AddUserRequest) (*models.User, error) {
	var existingUser models.User
	if database.GetDB().Where("discord_id = ?", body.DiscordId).First(&existingUser).Error == nil {
		return &existingUser, nil
	}

	return nil, nil
}

func updateUserEmail(user *models.User, email string) error {
	user.PaypalEmail = email
	if err := database.GetDB().Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// convertToInputModel converts a request body to a model insertable in DB
func convertToInputModel(body *pb_usr.AddUserRequest) models.User {
	return models.User{
		ID:          uuid.New(),
		Username:    body.Username,
		DiscordID:   body.DiscordId,
		PaypalEmail: body.PaypalEmail,
		CreatedAt:   time.Now(),
	}
}
