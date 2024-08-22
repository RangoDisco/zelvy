package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelvy/server/components"
	"github.com/rangodisco/zelvy/server/services"
	"github.com/rangodisco/zelvy/server/types"
)

func RegisterChartRoutes(r *gin.Engine) {
	r.GET("/charts", getCharts)
}

func getCharts(c *gin.Context) {
	var charts []types.Chart
	// Fetch radar chart (workout type)
	radar, err := services.GetWorkoutTypeChart()

	charts = append(charts, radar)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	components.Charts(charts).Render(c.Request.Context(), c.Writer)
}
