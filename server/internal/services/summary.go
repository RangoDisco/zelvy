package services

import (
	"fmt"
	"strconv"
	"time"

	"server/config/database"
	"server/internal/models"
	"server/pkg/types"

	"github.com/google/uuid"
)

func FetchSummaryByDate(date string) (models.Summary, error) {
	var summary models.Summary
	// Start building query
	q := database.GetDB().Preload("Workouts").Preload("Metrics").Preload("Winner").
		Order("date desc")

	// In case a date is provided, we want to fetch the summary for that date
	if date != "" {
		sod, eod, err := FormatDate(date)
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

func CreateSummaryViewModel(summary models.Summary) (types.SummaryViewModel, error) {

	// Build summary response
	var res types.SummaryViewModel
	res.ID = summary.ID.String()
	res.Date = fmt.Sprintf("%d %s %d", summary.Date.Day(), summary.Date.Month(), summary.Date.Year())
	res.Winner.DiscordID = summary.Winner.DiscordID

	// Compare metrics with goals to see wheter they are successful or not
	metrics, err := CompareMetricsWithGoals(summary.Metrics, summary.Workouts)
	if err != nil {
		return types.SummaryViewModel{}, err
	}

	res.Metrics = metrics

	// Add workouts to metrics object
	for _, w := range summary.Workouts {
		workout := ConvertToWorkoutViewModel(w)
		res.Workouts = append(res.Workouts, workout)
	}

	return res, nil
}

// Convert ms to hour and minute format
func ConvertMsToHour(ms float64) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

func FormatDate(stringDate string) (time.Time, time.Time, error) {
	// Parse date
	sod, err := time.Parse("2006-01-02", stringDate)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	eod := sod.Add(24 * time.Hour)

	return sod, eod, nil
}

func PickWinner() uuid.UUID {
	var u models.User

	// Get user from db randomly
	database.GetDB().Order("RANDOM()").First(&u)

	return u.ID
}
