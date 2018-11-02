package auth

import (
	"github.com/MateuszJonak/movies-api/models"
	"github.com/MateuszJonak/movies-api/storage"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	// to custom payload
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			models.IdentityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var payload models.User
	var getDB = storage.GetDB()

	if err := c.ShouldBind(&payload); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	user := models.User{}

	if err := getDB.Where(&models.User{Email: payload.Email}).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &models.User{
		Email: payload.Email,
	}, nil
}
