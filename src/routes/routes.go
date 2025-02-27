package routes

import (
	"JWT-Auth-Serviceject/src/controllers"
	"JWT-Auth-Serviceject/src/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Добро пожаловать в защищённый API!"})
	})

	return r
}
