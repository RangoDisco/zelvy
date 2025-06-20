package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/config/database"
	"server/pkg/types"
	"server/tests"
	"server/tests/factories"

	"github.com/stretchr/testify/assert"
)

func TestAddSummary(t *testing.T) {
	// Create an example input model
	summaryToCreate := factories.CreateSummaryInputModel()

	w := httptest.NewRecorder()

	userJson, _ := json.Marshal(summaryToCreate)

	req := SendRequest("POST", "/api/summaries", userJson)

	tests.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetSummary(t *testing.T) {
	// Insert summary in db
	summary := factories.CreateSummaryModel()
	res := database.GetDB().Create(&summary)
	assert.NoError(t, res.Error)

	// Try fetching the latest summary
	w := httptest.NewRecorder()
	req := SendRequest("GET", "/summaries", nil)

	tests.Router.ServeHTTP(w, req)

	// Ensure that the response is OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Ensure that the response is a JSON
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	// Ensure that the response is the correct summary
	var summaryResponse types.SummaryViewModel
	err := json.Unmarshal(w.Body.Bytes(), &summaryResponse)
	assert.NoError(t, err)
	assert.Equal(t, summary.ID.String(), summaryResponse.ID)
}
