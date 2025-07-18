package services

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"server/config/database"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

// FetchSummaryByDate retrieves summary from db by date
func FetchSummaryByDate(date string) (models.Summary, error) {
	var summary models.Summary
	// Start building the query
	q := database.GetDB().Preload("Workouts").Preload("Metrics").Preload("Metrics.Goal").Preload("Winner").
		Order("date desc")

	// In case a date is provided, we want to fetch the summary for that date
	if date != "" {
		sod, eod, err := formatDate(date)
		if err != nil {
			return models.Summary{}, err
		}

		// Add clause with the date provided
		q.Where("date >= ? AND date < ?", sod, eod)
	}

	// Query handlers from today
	if err := q.First(&summary).Error; err != nil {
		return models.Summary{}, err
	}

	return summary, nil
}

// CreateSummaryViewModel Converts a summary model to ViewModel that matches fields needed by the frontend
func CreateSummaryViewModel(summary *models.Summary) (types.SummaryViewModel, error) {
	var res types.SummaryViewModel

	var goals []models.Goal
	database.GetDB().Where("active = ?", true).Find(&goals)

	// Check if each goal has been fulfilled
	for _, g := range goals {
		idx := slices.IndexFunc(summary.Metrics, func(m models.Metric) bool {
			return m.GoalID == g.ID
		})

		var m *models.Metric

		// Workouts related goals don't have a related metric
		if idx == -1 {
			m = nil
		} else {
			m = &summary.Metrics[idx]
		}

		goalModel, err := convertToGoalViewModel(m, &g, &summary.Workouts)
		if err != nil {
			return types.SummaryViewModel{}, err
		}
		res.Goals = append(res.Goals, goalModel)
	}

	// Add workouts to the metrics object
	for _, w := range summary.Workouts {
		workout := ConvertToWorkoutViewModel(&w)
		res.Workouts = append(res.Workouts, workout)
	}

	res.ID = summary.ID.String()
	res.Date = fmt.Sprintf("%d %s %d", summary.Date.Day(), summary.Date.Month(), summary.Date.Year())
	res.Winner.DiscordID = summary.Winner.DiscordID

	return res, nil
}

// convertMsToHour and minute format
func convertMsToHour(ms float64) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

func formatDate(stringDate string) (time.Time, time.Time, error) {
	// Parse date
	sod, err := time.Parse("2006-01-02", stringDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	eod := sod.Add(24 * time.Hour)

	return sod, eod, nil
}

// PickWinner randomly selects a user from the database
func PickWinner() (uuid.UUID, error) {
	var u models.User

	// Get user from db randomly
	if err := database.GetDB().Order("RANDOM()").First(&u); err.Error != nil {
		return uuid.UUID{}, err.Error
	}

	return u.ID, nil
}
