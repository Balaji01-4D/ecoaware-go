package controllers

import (
	"net/http"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {

	var categories []models.Category

	err := initializer.DB.Find(&categories).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to fetch the categories",
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}