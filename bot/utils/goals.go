package utils

import (
	"fmt"
)

type RequestBody struct {
	Goals []string `json:"goals"`
}

/**
 * Send an array of goal type to be disabled for today
 */
func SetOffDay(goals []string) (bool, error) {
	var b = RequestBody{
		goals,
	}

	resp, err := Request("POST", "/api/offdays", b)

	if err != nil {
		return false, fmt.Errorf("error sending goals to disable: %v", err)
	}

	if resp.StatusCode() != 200 {
		return false, fmt.Errorf("error disabling goals: %v", resp.Status)
	}

	return true, nil
}
