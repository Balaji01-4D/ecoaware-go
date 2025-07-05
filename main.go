package main

import (
	"time"

	"github.com/Balaji01-4D/ecoware-go/controllers"
	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is home",
		})
	})
	r.POST("/auth/register", controllers.RegisterUser)
	r.POST("/auth/login", controllers.Login)
	r.GET("/auth/me", controllers.Me)
	r.GET("/categories", controllers.GetAllCategories)

	protected := r.Group("/")
	protected.Use(middleware.RequireAuth())

	protected.GET("/complaints", controllers.GetMyComplaints)
	protected.POST("/complaints", controllers.AddComplaints)
	protected.GET("/complaints/:id", controllers.GetComplaintByID)
	protected.PUT("/user/password", controllers.UpdatePassword)
	protected.PUT("/user/profile", controllers.UpdateUserByUser)


	admin := r.Group("/admin")
	admin.Use(middleware.RequireAuth(), middleware.RequireAdmin())

	admin.GET("/users", controllers.GetAllUsers)
	admin.DELETE("/users/:id", controllers.DeleteUser)
	admin.PUT("/users/:id", controllers.UpdateUser)
	admin.GET("/complaints", controllers.GetAllComplaints)
	admin.PUT("/complaints/:id/status", controllers.UpdateComplaintStatus)


	r.Run()
}
