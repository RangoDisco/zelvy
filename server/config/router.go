package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/components"
	"github.com/rangodisco/zelby/server/gintemplrenderer"
	"github.com/rangodisco/zelby/server/handlers"
	"github.com/rangodisco/zelby/server/middlewares"
)

func SetupRouter() *gin.Engine {
	// Start gin server
	r := gin.Default()

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Middleware to check API key in header
	r.Use(middlewares.CheckKey([]string{"/", "/summaries", "/charts"}))

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
