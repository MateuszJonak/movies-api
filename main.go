package main

import (
	"log"

	"github.com/MateuszJonak/movies-api/handlers"
	"github.com/MateuszJonak/movies-api/middlewares"
	"github.com/MateuszJonak/movies-api/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	storage.Open()

	r := gin.Default()

	authMiddleware, err := middlewares.CreateJWTMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/signup", handlers.SignUp)

	r.NoRoute(authMiddleware.MiddlewareFunc(), handlers.Error404)

	auth := r.Group("/auth")

	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", handlers.Hello)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

	defer storage.CloseDB()
}
