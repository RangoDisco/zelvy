package factories

import (
	"time"

	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"
	"github.com/rangodisco/zelvy/server/internal/models"

	"github.com/google/uuid"
)

func CreateSummaryModel() *models.Summary {
	id := uuid.New()

	return &models.Summary{
		ID:       id,
		Date:     time.Now(),
		Metrics:  CreateMetricModels(id),
		Workouts: CreateWorkoutModels(id),
		Winner:   CreateWinner(),
	}
}

func CreateSummaryViewModel() pb_sum.GetSummaryResponse {
	return pb_sum.GetSummaryResponse{
		Id:       "id",
		Day:      "2024-01-01",
		Goals:    CreateGoalViewModels(),
		Workouts: CreateWorkoutViewModels(),
		Winner:   CreateWinnerViewModel(),
	}
}

func CreateSummaryInputModel() *pb_sum.AddSummaryRequest {
	return &pb_sum.AddSummaryRequest{
		Metrics:  CreateMetricInputModels(),
		Workouts: CreateWorkoutInputModels(),
	}

}

func CreateWinnerViewModel() *pb_usr.GetSummaryUserResponse {
	return &pb_usr.GetSummaryUserResponse{
		DiscordId: "1231231231",
	}
}

func CreateWinner() models.User {
	return models.User{
		DiscordID:   "123456789",
		Username:    "Test User",
		PaypalEmail: "dummy@test.com",
	}
}
