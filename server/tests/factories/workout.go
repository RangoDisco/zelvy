package factories

import (
	"github.com/google/uuid"
	pb_wkr "github.com/rangodisco/zelvy/gen/zelvy/workout"
	"github.com/rangodisco/zelvy/server/internal/models"
	"github.com/rangodisco/zelvy/server/internal/services"
)

func CreateWorkoutModels(summaryId uuid.UUID) []models.Workout {
	return []models.Workout{
		{
			ID:           uuid.New(),
			Name:         "Push 1",
			KcalBurned:   320,
			ActivityType: pb_wkr.WorkoutActivityType_STRENGTH.String(),
			Duration:     3600,
			SummaryID:    summaryId,
		}, {
			ID:           uuid.New(),
			Name:         "Marche",
			KcalBurned:   300,
			ActivityType: pb_wkr.WorkoutActivityType_WALK.String(),
			Duration:     3600,
			SummaryID:    summaryId,
		},
	}
}

func CreateWorkoutViewModels() []*pb_wkr.WorkoutViewModel {
	var workoutsViewModels []*pb_wkr.WorkoutViewModel

	workouts := CreateWorkoutModels(uuid.New())

	for _, w := range workouts {
		workoutsViewModels = append(workoutsViewModels, services.ConvertToWorkoutViewModel(&w))
	}

	return workoutsViewModels

}

func CreateWorkoutInputModels() []*pb_wkr.WorkoutInputModel {
	pushName := "Push 1"
	return []*pb_wkr.WorkoutInputModel{
		{
			Name:         &pushName,
			Duration:     3600,
			ActivityType: pb_wkr.WorkoutActivityType_STRENGTH,
			KcalBurned:   320,
		},
		{
			Duration:     3600,
			ActivityType: pb_wkr.WorkoutActivityType_WALK,
			KcalBurned:   300,
		},
	}
}
