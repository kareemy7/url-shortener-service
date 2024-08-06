package main

import (
	"url-shortener3/controllers"
	"url-shortener3/initializers"
	"url-shortener3/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SynceDB()
}
func main() {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/shorten", middleware.RequireAuth, controllers.Shorten)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/:shortCode", controllers.Retrieve)
	r.Run()
}
