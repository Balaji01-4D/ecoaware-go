package controllers

import (
	"net/http"
	"strconv"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	var user models.User
	
	if err := c.BindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}


	result := initializer.DB.Create(&user)

	if (result.Error != nil) {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {

	var users []models.User

	initializer.DB.Find(&users)

	c.JSON(200, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := getById(intId)
	if user != nil {
		c.JSON(http.StatusOK, user)
		return
	}
	c.Status(http.StatusNotFound)
}


func getById(id int) (*models.User){
	var user models.User
	result := initializer.DB.First(&user, id)

	if result.Error != nil {
		return nil
	}
	return  &user
} 


func UpdateUser(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User
	if err := initializer.DB.First(&user, id).Error; err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	if err := initializer.DB.Save(&user).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	if err := initializer.DB.Delete(&models.User{}, idParam).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)

}
