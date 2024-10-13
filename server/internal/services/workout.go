package services

import (
	"time"

	"server/config/database"
	"server/internal/enums"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
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

// FetchWorkoutsByDateRange fetches all the workouts in the database between two dates
func FetchWorkoutsByDateRange(startDate time.Time, endDate time.Time) ([]models.Workout, error) {
	var workouts []models.Workout
	err := database.GetDB().Raw("SELECT * FROM workouts w INNER JOIN summaries s ON w.summary_id = s.id WHERE s.date >= ? AND s.date < ?", startDate, endDate).Scan(&workouts).Error
	return workouts, err
}

// ConvertToWorkoutModel converts a WorkoutInputModel to a Workout model (used when creating a new workout)
func ConvertToWorkoutModel(w *types.WorkoutInputModel, summaryId uuid.UUID) models.Workout {
	return models.Workout{
		ID:           uuid.New(),
		SummaryID:    summaryId,
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         getWorkoutName(w),
		Duration:     w.Duration,
	}
}

// ConvertToWorkoutViewModel converts a Workout model to a WorkoutViewModel (used in the front)
func ConvertToWorkoutViewModel(w *models.Workout) types.WorkoutViewModel {
	return types.WorkoutViewModel{
		ID:           w.ID.String(),
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         w.Name,
		Duration:     ConvertMsToHour(w.Duration),
		Picto:        getWorkoutPicto(w.ActivityType),
	}
}

// CalculateMainWorkoutDuration sums the duration of all the gym sessions
func CalculateMainWorkoutDuration(workouts *[]models.Workout) float64 {
	var duration float64
	for _, w := range *workouts {
		if w.ActivityType == "strength" {
			duration = duration + w.Duration
		}
	}
	return duration
}

// CalculateExtraWorkoutDuration sums the duration of all cardio sessions
func CalculateExtraWorkoutDuration(workouts *[]models.Workout) float64 {
	var duration float64
	for _, w := range *workouts {
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
func getWorkoutName(w *types.WorkoutInputModel) string {
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
