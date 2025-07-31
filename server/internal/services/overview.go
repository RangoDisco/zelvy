package services

import (
	"server/config/database"
	"server/pkg/types"
	"time"
)

func getDateRangeFromString(sd string, ed string) (time.Time, time.Time, error) {
	fsd, err := time.Parse("2006-01-02", sd)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	fed, err := time.Parse("2006-01-02", ed)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return fsd, fed, nil
}

// GetWinnersList fetches all winners between two dates and order them by number of wins
func GetWinnersList(sd, ed string, limit int64) ([]types.WinnerViewModel, error) {
	fsd, fed, err := getDateRangeFromString(sd, ed)
	if err != nil {
		return []types.WinnerViewModel{}, err
	}

	var winners []types.WinnerViewModel
	err = database.GetDB().Raw("SELECT username as username, count(s.id) as wins FROM users u INNER JOIN summaries s ON s.winner_id = u.id WHERE s.date >= ? AND s.date <? GROUP BY username ORDER BY wins DESC LIMIT ?", fsd, fed, limit).Scan(&winners).Error
	if err != nil {
		return []types.WinnerViewModel{}, err
	}
	return winners, nil
}
