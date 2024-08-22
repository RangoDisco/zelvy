package factories

import (
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/services"
	"github.com/rangodisco/zelby/server/types"
)

func CreateWorkoutModels(summaryId uuid.UUID) []models.Workout {
	return []models.Workout{
		{
			ID:           uuid.New(),
			Name:         "Push 1",
			KcalBurned:   320,
			ActivityType: enums.WorkoutTypeStrength,
			Duration:     3600,
			SummaryID:    summaryId,
		}, {
			ID:           uuid.New(),
			Name:         "Marche",
			KcalBurned:   300,
			ActivityType: enums.WorkoutTypeWalking,
			Duration:     3600,
			SummaryID:    summaryId,
		},
	}
}

func CreateWorkoutViewModels() []types.WorkoutViewModel {
	var workoutsViewModels []types.WorkoutViewModel

	workouts := CreateWorkoutModels(uuid.New())

	for _, w := range workouts {
		workoutsViewModels = append(workoutsViewModels, services.ConvertToWorkoutViewModel(w))
	}

	return workoutsViewModels

}

func CreateWorkoutInputModels() []types.WorkoutInputModel {
	return []types.WorkoutInputModel{
		{
			Name:         "Push 1",
			Duration:     3600,
			ActivityType: enums.WorkoutTypeStrength,
			KcalBurned:   320,
		},
		{
			Name:         "Marche",
			Duration:     3600,
			ActivityType: enums.WorkoutTypeWalking,
			KcalBurned:   300,
		},
	}
}
