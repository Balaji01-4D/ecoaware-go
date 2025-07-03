package main

import (
	"github.com/Balaji01-4D/ecoware-go/controllers"
	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is home",
		})
	})
	router.POST("/user/register", controllers.RegisterUser)
	router.POST("/user/login", controllers.Login)

	
	api := router.Group("/api")
	api.Use(middleware.RequireAuth())

	api.GET("/user/validate", controllers.Validate)
	api.POST("/complaint", controllers.AddComplaints)
	api.PUT("/complaint/:id", controllers.UpdateComplaints)


	router.GET("/api/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.Run()
}
