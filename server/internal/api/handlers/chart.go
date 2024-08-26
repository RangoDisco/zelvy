package handlers

import (
	"server/components"
	"server/internal/services"
	"server/pkg/types"

	"github.com/gin-gonic/gin"
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
