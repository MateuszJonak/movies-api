package handlers

import (
	"github.com/MateuszJonak/movies-api/models"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	user, _ := c.Get(models.IdentityKey)
	c.JSON(200, gin.H{
		"userName": user.(*models.User).Email,
		"text":     "Hello World.",
	})
}
