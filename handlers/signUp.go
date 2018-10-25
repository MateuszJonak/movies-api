package handlers

import (
	"fmt"
	"net/http"

	"github.com/MateuszJonak/movies-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var payload models.UserSignUp
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 8)

	// check stored pass
	// if err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password)); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"userName":       payload.Username,
		"hashedPassword": hashedPassword,
	})
}
