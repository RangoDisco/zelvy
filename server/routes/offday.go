package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"time"
)

type CreateOffDayBody struct {
	Goals []uuid.UUID `json:"goals"`
}

func RegisterOffDayRoutes(r *gin.Engine) {

}

func setOffDay(c *gin.Context) {
	var body CreateOffDayBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, goal := range body.Goals {
		// Ensure that goal exists
		if err := database.DB.First(&models.Goal{}, "id = ?", goal).Error; err != nil {
			continue
		}

		// Create off day for this specific goal
		offday := models.Offday{
			GoalID: goal,
			Day:    time.Now(),
		}

		if err := database.DB.Create(&offday).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

	}
}
