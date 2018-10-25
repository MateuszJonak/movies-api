package handlers

import (
	"log"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func Error404(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
