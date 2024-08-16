package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/enums"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
)

func FetchChartWorkouts() ([]models.Workout, []models.Workout, error) {
	// Get dates of the current week
	sow := time.Now().Add(-168 * time.Hour)
	eow := time.Now()

	// Get the number of workouts between 7 days ago and now
	thisWeek, err := FetchWorkoutsByDateRange(sow, eow)
	if err != nil {
		return []models.Workout{}, []models.Workout{}, err
	}

	// Get the number of workouts between 14 days ago and 7 days ago
	lastWeek, err := FetchWorkoutsByDateRange(sow.Add(-168*time.Hour), sow)
	if err != nil {
		return []models.Workout{}, []models.Workout{}, err
	}

	return thisWeek, lastWeek, nil

}

func FetchWorkoutsByDateRange(startDate time.Time, endDate time.Time) ([]models.Workout, error) {
	var workouts []models.Workout
	err := database.DB.Raw("SELECT * FROM workouts w INNER JOIN summaries s ON w.summary_id = s.id WHERE s.date >= ? AND s.date < ?", startDate, endDate).Scan(&workouts).Error
	return workouts, err
}

func ConvertToWorkoutModel(w types.WorkoutInputModel, summaryId uuid.UUID) models.Workout {
	return models.Workout{
		ID:           uuid.New(),
		SummaryID:    summaryId,
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         getWorkoutName(w),
		Duration:     w.Duration,
	}
}

func ConvertToWorkoutViewModel(w models.Workout) types.WorkoutViewModel {
	return types.WorkoutViewModel{
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
		return "/assets/img/cardio.png"
	}
}

// Handles name based on activity type in case null
func getWorkoutName(w types.WorkoutInputModel) string {
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
