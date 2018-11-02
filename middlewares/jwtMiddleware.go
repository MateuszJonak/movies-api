package middlewares

import (
	"os"
	"time"

	"github.com/MateuszJonak/movies-api/auth"
	"github.com/MateuszJonak/movies-api/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func CreateJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	pwd, _ := os.Getwd()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour * 24,
		IdentityKey:      models.IdentityKey,
		SigningAlgorithm: "RS256",
		PrivKeyFile:      pwd + "/middlewares/jwtRS256.key",
		PubKeyFile:       pwd + "/middlewares/jwtRS256.key.pub",
		PayloadFunc:      auth.PayloadFunc,
		Authenticator:    auth.Authenticator,
		IdentityHandler:  auth.IdentityHandler,
		Authorizator:     auth.Authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	return authMiddleware, err
}
