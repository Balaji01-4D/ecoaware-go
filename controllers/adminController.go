package controllers

import (
	"net/http"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	user := c.MustGet("user").(models.User)

	if user.Role != models.RoleAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "permission denied",
		})
		return
	}

	var users []models.User

	err := initializer.DB.Find(&users).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, users)

}
