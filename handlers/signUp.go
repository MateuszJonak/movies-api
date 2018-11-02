package handlers

import (
	"net/http"

	"github.com/MateuszJonak/movies-api/models"
	"github.com/MateuszJonak/movies-api/storage"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var payload models.User
	var getDB = storage.GetDB()

	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{}
	copier.Copy(&user, &payload)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 8)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	if err := getDB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"email": payload.Email,
	})
}
