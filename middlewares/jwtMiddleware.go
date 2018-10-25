package middlewares

import (
	"os"
	"time"

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
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// to custom payload
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					models.IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// to get user and use it to authorizator
			claims := jwt.ExtractClaims(c)
			return &models.User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var payload models.UserSignIn
			if err := c.ShouldBind(&payload); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := payload.Username
			password := payload.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &models.User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	return authMiddleware, err
}
