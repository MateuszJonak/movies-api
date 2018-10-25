package handlers

import (
	"github.com/MateuszJonak/movies-api/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(models.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*models.User).UserName,
		"text":     "Hello World.",
	})
}
