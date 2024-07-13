package routes

import (
	"github.com/gin-gonic/gin"
	"rangodisco.eu/zelby-server/database"
	"rangodisco.eu/zelby-server/models"
)

type CreateGoalBody struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

func RegisterGoalRoutes(r *gin.Engine) {
	r.GET("/api/goals", getGoals)
	r.POST("/api/goals", addGoal)
}

func getGoals(c *gin.Context) {
	var goals []models.Goal

	if err := database.DB.Find(&goals).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, goals)
}

func addGoal(c *gin.Context) {
	// Parse body
	var body CreateGoalBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Convert to model
	goal := models.Goal{
		Type:  body.Type,
		Value: body.Value,
	}

	if err := database.DB.Create(&goal).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, goal)

}
