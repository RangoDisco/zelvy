package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rangodisco/zelby/server/tests"
	"github.com/rangodisco/zelby/server/tests/factories"

	"github.com/stretchr/testify/assert"
)

func TestAddSummary(t *testing.T) {
	// Setup router

	// Create example input model
	summaryToCreate := factories.CreateSummaryInputModel()

	w := httptest.NewRecorder()

	userJson, _ := json.Marshal(summaryToCreate)

	req := SendRequest("POST", "/api/summaries", userJson)

	tests.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}
