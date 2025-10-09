package services

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/google/uuid"
	pb_goa "github.com/rangodisco/zelvy/gen/zelvy/goal"
	pb_wrk "github.com/rangodisco/zelvy/gen/zelvy/workout"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/enums"
	"github.com/rangodisco/zelvy/server/internal/models"
)

// ConvertToWorkoutModel converts a WorkoutInputModel to a Workout model (used when creating a new workout)
func ConvertToWorkoutModel(w *pb_wrk.WorkoutInputModel, summaryId uuid.UUID) models.Workout {
	return models.Workout{
		ID:           uuid.New(),
		SummaryID:    summaryId,
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType.String(),
		Name:         getWorkoutName(w),
		Duration:     w.Duration,
		DoneAt:       w.DoneAt.AsTime(),
	}
}

// ConvertToWorkoutViewModel used when fetching a summary, converts a Workout model to a WorkoutViewModel
func ConvertToWorkoutViewModel(w *models.Workout) *pb_wrk.WorkoutViewModel {
	return &pb_wrk.WorkoutViewModel{
		Id:           w.ID.String(),
		KcalBurned:   w.KcalBurned,
		ActivityType: w.ActivityType,
		Name:         w.Name,
		Duration:     convertMsToHour(w.Duration),
		Picto:        getWorkoutPicto(w.ActivityType),
		DoneAt:       timestamppb.New(w.DoneAt),
	}
}

// calculateWorkoutDuration used when fetching a summary, add the duration of all the workout from a given type
func calculateWorkoutDuration(w *[]models.Workout, target string) float64 {
	var duration float64
	for _, w := range *w {
		if (target == pb_goa.GoalType_MAIN_WORKOUT_DURATION.String() && w.ActivityType == pb_wrk.WorkoutActivityType_STRENGTH.String()) ||
			(target == pb_goa.GoalType_EXTRA_WORKOUT_DURATION.String() && w.ActivityType != pb_wrk.WorkoutActivityType_STRENGTH.String()) {
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
func getWorkoutName(w *pb_wrk.WorkoutInputModel) string {
	if w.Name != nil {
		return *w.Name
	}

	switch w.ActivityType {
	case pb_wrk.WorkoutActivityType_STRENGTH:
		return "Gym"
	case pb_wrk.WorkoutActivityType_RUNNING:
		return "Running"
	case pb_wrk.WorkoutActivityType_CYCLING:
		return "Cycling"
	case pb_wrk.WorkoutActivityType_WALK:
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
