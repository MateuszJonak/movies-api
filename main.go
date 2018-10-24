package main

import (
	"log"

	"github.com/MateuszJonak/movies-api/middlewares"
	"github.com/MateuszJonak/movies-api/models"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(models.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*models.User).UserName,
		"text":     "Hello World.",
	})
}

func main() {
	r := gin.Default()

	authMiddleware, err := middlewares.CreateJWTMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
