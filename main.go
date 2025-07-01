package main

import (
	"github.com/Balaji01-4D/ecoware-go/controllers"
	"github.com/Balaji01-4D/ecoware-go/initializer"
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
			"message" : "this is home",
		})
	})
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.AddUser)
	router.GET("/user/:id", controllers.GetUserById)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.Run()
}