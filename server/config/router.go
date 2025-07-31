package config

import (
	"net/http"

	"server/components"
	"server/internal/api/handlers"
	"server/internal/api/middlewares"
	"server/pkg/gintemplrenderer"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Start gin server
	r := gin.Default()

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Defines public routes (mainly the dashboard's ones)
	r.Use(middlewares.CheckKey([]string{"/", "/overview", "/overview/winners", "/charts"}))

	// Register handlers from handlers package
	handlers.RegisterSummaryRoutes(r)
	handlers.RegisterGoalRoutes(r)
	handlers.RegisterOffDayRoutes(r)
	handlers.RegisterUserRoutes(r)
	handlers.RegisterChartRoutes(r)
	handlers.RegisterOverviewRoutes(r)

	// Serve static files
	r.Static("/assets", "assets")

	// Handle 404
	r.NoRoute(func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusNotFound, components.NotFound())
		c.Render(404, r)
	})

	return r
}
