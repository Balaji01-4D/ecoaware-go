package controllers

import (
	"net/http"
	"strconv"

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


func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	if err := initializer.DB.Delete(&models.User{}, idParam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)

}


func UpdateUser(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var input struct {
		Name  string	`json:"name"`
		Email string	`json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := initializer.DB.Model(&models.User{}).
	Where("id = ?", id).
	Update("name", input.Name).
	Update("email", input.Email).
	Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"updated successfully",
	})
}