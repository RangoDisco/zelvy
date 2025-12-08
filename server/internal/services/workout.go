package services

import (
	"errors"
	"fmt"
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
		DoneAt:       w.DoneAt.Format(time.RFC3339),
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

// FetchWorkoutsByDateRange fetches all the workouts in the database between two dates
func FetchWorkoutsByDateRange(startDate *string, endDate *string) ([]models.Workout, error) {
	sd, err := GetTimeFromString(startDate)
	if err != nil {
		return []models.Workout{}, errors.New(fmt.Sprintf("could not parse date: %s", err))
	}

	ed, err := GetTimeFromString(endDate)
	if err != nil {
		return []models.Workout{}, errors.New(fmt.Sprintf("could not parse date: %s", err))
	}

	var workouts []models.Workout
	err = database.GetDB().Raw("SELECT * FROM workouts w INNER JOIN summaries s ON w.summary_id = s.id WHERE s.date >= ? AND s.date < ?", sd, ed).Scan(&workouts).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not fetch workouts"))
	}

	return workouts, nil
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
