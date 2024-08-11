package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/types"
)

func RegisterChartRoutes(r *gin.Engine) {
	r.GET("/charts", getCharts)
}

func getCharts(c *gin.Context) {
	data := []types.Chart{
		{
			Type:   "radar",
			Labels: []string{"Salle", "Marche", "Footing"},
			Datasets: []types.Dataset{
				{
					Label: "Cette semaine",
					Data:  []int{6, 7, 2},
				},
				{
					Label: "La semaine dernière",
					Data:  []int{3, 4, 5},
				},
			},
		},
	}

	components.Charts(data).Render(c.Request.Context(), c.Writer)
}
