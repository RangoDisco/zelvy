package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/gintemplrenderer"
	"github.com/rangodisco/zelby/server/handlers"
	"github.com/rangodisco/zelby/server/tests/factories"

	"github.com/stretchr/testify/assert"
)

func TestAddSummary(t *testing.T) {
	// Setup router
	r := SetupRouter()

	// Create example input model
	summaryToCreate := factories.CreateSummaryInputModel()

	w := httptest.NewRecorder()

	userJson, _ := json.Marshal(summaryToCreate)
	req, _ := http.NewRequest("POST", "/api/summaries", strings.NewReader(string(userJson)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

func SetupRouter() *gin.Engine {
	// Start gin server
	r := gin.Default()

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Middleware to check API key in header
	//r.Use(middlewares.CheckKey())

	// Register handlers from handlers package
	handlers.RegisterSummaryRoutes(r)
	handlers.RegisterGoalRoutes(r)
	handlers.RegisterOffDayRoutes(r)
	handlers.RegisterUserRoutes(r)
	handlers.RegisterChartRoutes(r)

	// Serve static files
	r.Static("/assets", "./assets")

	// Handle 404
	r.NoRoute(func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusNotFound, components.NotFound())
		c.Render(404, r)
	})

	return r
}
