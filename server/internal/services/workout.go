package services

import (
	"time"

	pb_wrk "github.com/rangodisco/zelvy/gen/zelvy/workout"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/enums"
	"github.com/rangodisco/zelvy/server/internal/models"
	"github.com/rangodisco/zelvy/server/pkg/types"

	"github.com/google/uuid"
)

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

// ConvertToWorkoutViewModel used when fetching a summary, converts a Workout model to a WorkoutViewModel
func ConvertToWorkoutViewModel(w *models.Workout) pb_wrk.WorkoutViewModel {
	return pb_wrk.WorkoutViewModel{
		Id: w.ID.String(),
		// TODO: fix
		KcalBurned:   int64(w.KcalBurned),
		ActivityType: w.ActivityType,
		Name:         w.Name,
		Duration:     convertMsToHour(w.Duration),
		Picto:        getWorkoutPicto(w.ActivityType),
	}
}

// calculateWorkoutDuration used when fetching a summary, add the duration of all the workout from a given type
func calculateWorkoutDuration(w *[]models.Workout, target string) float64 {
	var duration float64
	for _, w := range *w {
		if (target == enums.MainWorkoutDuration && w.ActivityType == "strength") || (target == enums.ExtraWorkoutDuration && w.ActivityType != "strength") {
			duration = duration + w.Duration
		}
	}

	return duration
}

// fetchWorkoutsByDateRange fetches all the workouts in the database between two dates
func fetchWorkoutsByDateRange(startDate time.Time, endDate time.Time) ([]models.Workout, error) {
	var workouts []models.Workout
	err := database.GetDB().Raw("SELECT * FROM workouts w INNER JOIN summaries s ON w.summary_id = s.id WHERE s.date >= ? AND s.date < ?", startDate, endDate).Scan(&workouts).Error
	return workouts, err
}

func getWorkoutPicto(activityType string) string {
	switch activityType {
	case enums.WorkoutTypeStrength:
		return "/assets/img/strength.png"
	default:
		return "/assets/img/cardio.png"
	}
}

// Handles name based on the activity's type in case null
func getWorkoutName(w *types.WorkoutInputModel) string {
	if w.Name != "" {
		return w.Name
	}

	switch w.ActivityType {
	case enums.WorkoutTypeStrength:
		return "Gym"
	case enums.WorkoutTypeRunning:
		return "Running"
	case enums.WorkoutTypeCycling:
		return "Cycling"
	case enums.WorkoutTypeWalking:
		return "Walking"
	default:
		return "Random silly workout"
	}
}

func fetchChartWorkouts() ([]models.Workout, []models.Workout, error) {
	// Get dates of the current week
	sow := time.Now().Add(-168 * time.Hour)
	eow := time.Now()

	// Get the number of workouts during the last 7 days
	thisWeek, err := fetchWorkoutsByDateRange(sow, eow)
	if err != nil {
		return []models.Workout{}, []models.Workout{}, err
	}

	// Get the number of workouts between 14 days ago and 7 days ago
	lastWeek, err := fetchWorkoutsByDateRange(sow.Add(-168*time.Hour), sow)
	if err != nil {
		return []models.Workout{}, []models.Workout{}, err
	}

	return thisWeek, lastWeek, nil

}
