package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
)

type CreateGoalBody struct {
	Type       string  `json:"type"`
	Name       string  `json:"name"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	Comparison string  `json:"comparison"`
}

func RegisterGoalRoutes(r *gin.Engine) {
	r.GET("/api/goals", getGoals)
	r.POST("/api/goals", addGoal)
}

func getGoals(c *gin.Context) {
	var goals []models.Goal

	if err := database.GetDB().Find(&goals).Error; err != nil {
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
		Type:       body.Type,
		Value:      body.Value,
		Name:       body.Name,
		Unit:       body.Unit,
		Comparison: body.Comparison,
	}

	if err := database.GetDB().Create(&goal).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, goal)

}
