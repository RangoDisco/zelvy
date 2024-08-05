package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rangodisco/zelby/server/database"
	"github.com/rangodisco/zelby/server/models"
	"github.com/rangodisco/zelby/server/types"
	"net/http"
	"time"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/api/users", addUser)
}

// ROUTES
func addUser(c *gin.Context) {
	var body types.UserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert to model
	u := models.User{
		ID:          uuid.New(),
		Username:    body.Username,
		DiscordID:   body.DiscordID,
		PaypalEmail: body.PaypalEmail,
		CreatedAt:   time.Now(),
	}

	// Persist
	if err := database.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User created")
}
