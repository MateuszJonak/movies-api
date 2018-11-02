package auth

import (
	"github.com/MateuszJonak/movies-api/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func IdentityHandler(c *gin.Context) interface{} {
	// to get user and use it to authorizator
	claims := jwt.ExtractClaims(c)
	return &models.User{
		Email: claims["id"].(string),
	}
}
func Authorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*models.User); ok {
		return true
	}

	return false
}
