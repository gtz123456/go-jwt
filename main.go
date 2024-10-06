package main

import (
	"jwt/controllers"
	"jwt/db"
	middleware "jwt/middleware"
	"jwt/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
	db.Connect()
	db.Sync()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
