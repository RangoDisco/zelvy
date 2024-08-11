package utils

import (
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
)

func ConvertToWorkoutModel(w types.WorkoutData, summaryId uuid.UUID) models.Workout {
	return models.Workout{
		ID:           uuid.New(),
		SummaryID:    summaryId,
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         getWorkoutName(w),
		Duration:     w.Duration,
	}
}

func ConvertToWorkoutResponse(w models.Workout) types.WorkoutResponse {
	return types.WorkoutResponse{
		ID:           w.ID.String(),
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         w.Name,
		Duration:     ConvertMsToHour(w.Duration),
		Picto:        getWorkoutPicto(w.ActivityType),
	}
}

// Calculate main workout duration
func CalculateMainWorkoutDuration(workouts []models.Workout) float64 {
	var duration float64
	for _, w := range workouts {
		if w.ActivityType == "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

// Calculate extra workout duration
func CalculateExtraWorkoutDuration(workouts []models.Workout) float64 {
	var duration float64
	for _, w := range workouts {
		if w.ActivityType != "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

func getWorkoutPicto(activityType string) string {
	switch activityType {
	case enums.WorkoutTypeStrength:
		return "/assets/img/strength.png"
	default:
		return "/assets/img/strength.png"
	}
}

// Handles name based on activity type in case null
func getWorkoutName(w types.WorkoutData) string {
	if w.Name != "" {
		return w.Name
	}

	switch w.ActivityType {
	case enums.WorkoutTypeStrength:
		return "Séance de musculation"
	case enums.WorkoutTypeRunning:
		return "Footing"
	case enums.WorkoutTypeCycling:
		return "Vélo"
	case enums.WorkoutTypeWalking:
		return "Marche"
	default:
		return "Séance de sport"
	}
}
