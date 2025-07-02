package main

import (
	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.User{})
	initializer.DB.AutoMigrate(&models.Complaint{})
}