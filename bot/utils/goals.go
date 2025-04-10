package utils

import (
	"fmt"
	"github.com/rangodisco/zelvy/bot/types"
)

// SetOffDay disables any given goal for today
func SetOffDay(goals []string) (bool, error) {
	var b = types.GoalRequestBody{
		Goals: goals,
	}

	resp, err := Request("POST", "/api/offdays", b)

	if err != nil {
		return false, fmt.Errorf("error sending goals to disable: %v", err)
	}

	if resp.StatusCode() != 200 {
		return false, fmt.Errorf("error disabling goals: %v", resp.StatusCode())
	}

	return true, nil
}
