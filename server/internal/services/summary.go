package services

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"time"

	pb_sum "github.com/rangodisco/zelvy/gen/zelvy/summary"
	pb_usr "github.com/rangodisco/zelvy/gen/zelvy/user"

	"github.com/google/uuid"
	"github.com/rangodisco/zelvy/server/config/database"
	"github.com/rangodisco/zelvy/server/internal/models"
)

// FetchSummaryByDate retrieves summary from db by date
func FetchSummaryByDate(date string) (models.Summary, error) {
	var s models.Summary
	// Start building the query
	q := database.GetDB().Preload("Workouts").Preload("Metrics").Preload("Metrics.Goal").Preload("Winner").
		Order("date desc").Where("deleted_at IS null")

	// In case a date is provided, we want to fetch the summary for that date
	if date != "" {
		sod, err := GetTimeFromString(&date)
		if err != nil {
			return models.Summary{}, err
		}
		eod := sod.Add(24 * time.Hour)

		// Add clause with the date provided
		q.Where("date >= ? AND date < ?", sod, eod)
	}

	// Query handlers from given day
	if err := q.First(&s).Error; err != nil {
		return models.Summary{}, err
	}

	return s, nil
}

// FindHeatmapResults fetches each day from a calendar between two dates and returns the summary for each
func FindHeatmapResults(startDate, endDate string) ([]*pb_sum.HeatmapItemViewModel, error) {
	parsedStartDate, err := GetTimeFromString(&startDate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse start date: %s", err))
	}

	parsedEndDate, err := GetTimeFromString(&endDate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not parse end date: %s", err))
	}

	if parsedStartDate.After(parsedEndDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	var items []*pb_sum.HeatmapItemViewModel
	err = database.GetDB().
		Table("calendar").
		Select("summaries.id AS id, COALESCE(summaries.success, true) AS success, calendar.date_without_time AS date").
		Joins("LEFT JOIN summaries on DATE(calendar.date_without_time) = DATE(summaries.date)").
		Where("summaries.deleted_at IS null AND calendar.date_without_time >= ? AND calendar.date_without_time <= ?", parsedStartDate, parsedEndDate).
		Order("calendar.date_without_time ASC").
		Scan(&items).Error

	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not fetch heatmap items"))
	}

	return items, nil
}

// CreateSummaryViewModel Converts a summary model to ViewModel that matches fields needed by the frontend
func CreateSummaryViewModel(summary *models.Summary) (*pb_sum.GetSummaryResponse, error) {
	var res pb_sum.GetSummaryResponse

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

		goalModel, err := convertToGoalViewModel(m, &g)
		if err != nil {
			continue
		}
		res.Goals = append(res.Goals, &goalModel)
	}

	// Add workouts to the metrics object
	for _, w := range summary.Workouts {
		workout := ConvertToWorkoutViewModel(&w)
		res.Workouts = append(res.Workouts, workout)
	}

	res.Id = summary.ID.String()
	res.Day = fmt.Sprintf("%d %s %d", summary.Date.Day(), summary.Date.Month(), summary.Date.Year())

	winner := pb_usr.GetSummaryUserResponse{DiscordId: summary.Winner.DiscordID, Name: summary.Winner.Username}
	res.Winner = &winner

	return &res, nil
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

// convertMsToHour and minute format
func convertMsToHour(ms float64) string {
	duration := time.Duration(ms) * time.Second
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
}

// getDateFromString converts a string date into a time.Time and check that if it is valid
func getDateFromString(stringDate string) (time.Time, error) {
	parsedDate, err := time.Parse(time.DateOnly, stringDate)
	if err != nil {
		return time.Time{}, errors.New(fmt.Sprintf("could not parse date: %s", err))
	}

	return parsedDate, nil
}
