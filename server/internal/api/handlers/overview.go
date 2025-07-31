package handlers

import (
	"net/http"
	"server/components"
	"server/internal/services"
	"server/pkg/gintemplrenderer"
	"server/pkg/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterOverviewRoutes(r *gin.Engine) {
	r.GET("/overview", getOverview)
	r.GET("/overview/winners", getTopWinners)
}

func getOverview(c *gin.Context) {
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, components.Overview())
	c.Render(http.StatusOK, r)
}

func getTopWinners(c *gin.Context) {
	sd, _ := c.GetQuery("sd")
	ed, _ := c.GetQuery("ed")
	winners, err := services.GetWinnersList(sd, ed, 1)
	if len(winners) == 0 || err != nil {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, components.NotFound())
		c.Render(http.StatusOK, r)
	}

	stat := types.StatViewModel{
		Picto:       "ancd",
		Title:       winners[0].Username,
		Description: "Most wins",
		Value:       strconv.Itoa(winners[0].Wins),
	}

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, components.BentoStat(stat))
	c.Render(http.StatusOK, r)
}

//func getCharts(c *gin.Context) {
//	var charts []types.Chart
//	// Fetch radar chart (workout type)
//	radar, err := services.GetWorkoutTypeChart()
//
//	charts = append(charts, radar)
//
//	if err != nil {
//		c.JSON(500, gin.H{"error": err.Error()})
//		return
//	}
//
//	components.Charts(charts).Render(c.Request.Context(), c.Writer)
//}
