package services

import (
	"errors"
	"fmt"
	"time"
)

// GetTimeFromString converts a string date into a time.Time and check that if it is valid
func GetTimeFromString(stringDate *string) (time.Time, error) {
	var date time.Time
	var err error

	if stringDate == nil {
		date = time.Now()
		return date, nil
	}

	parsedDate, err := time.Parse(time.DateOnly, *stringDate)
	if err != nil {
		return time.Time{}, errors.New(fmt.Sprintf("could not parse date: %s", err))
	}

	return parsedDate, nil
}
