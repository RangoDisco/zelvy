package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"time"
)

type CreateOffDayBody struct {
	Goals []string `json:"goals"`
}

func RegisterOffDayRoutes(r *gin.Engine) {
	r.POST("/api/offdays", setOffDay)
}

func setOffDay(c *gin.Context) {
	var body CreateOffDayBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, og := range body.Goals {
		var goal models.Goal

		// Ensure that goal exists
		if err := database.DB.First(&goal, "type = ?", og).Error; err != nil {
			continue
		}

		// Create off day for this specific goal
		offday := models.Offday{
			GoalID: goal.ID,
			Day:    time.Now(),
		}

		if err := database.DB.Create(&offday).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Offday created"})
	}
}
