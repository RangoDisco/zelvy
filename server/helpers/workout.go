package helpers

import (
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
)

func PopulateWorkoutMetric(duration int, goalValue int, name string, shouldThresholdBeSmaller bool) types.Metric {
	return types.Metric{
		Value:        duration,
		DisplayValue: ConvertMsToHour(duration),
		Threshold:    ConvertMsToHour(goalValue),
		Name:         name,
		Success:      IsMetricSuccessful(duration, goalValue, shouldThresholdBeSmaller),
	}
}

// Handles name based on activity type in case null
func GetWorkoutName(w types.WorkoutData) string {
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

func ConvertToWorkoutModel(w types.WorkoutData, metricRef uuid.UUID) models.Workout {
	return models.Workout{
		ID:           uuid.New(),
		MetricsRefer: metricRef,
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         GetWorkoutName(w),
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
	}
}
