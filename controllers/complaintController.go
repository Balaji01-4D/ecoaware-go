package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)


func AddComplaints(c *gin.Context) {
	user, exist := c.Get("user")

	if ! exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"not user found",
		})
		return
	}

	authUser := user.(models.User)

	var body struct {
		Title string		
		Description string	
		ImagePath string
	}


	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"failed to read body",
		})
		return
	}

	complaint := models.Complaint{
		Title: body.Title,
		Description: body.Description,
		ImagePath: body.ImagePath,
		CreatedBy: authUser.ID,
		CreatedAt: time.Now(),
	}

	
	result := initializer.DB.Create(&complaint)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"fail to add the complaint",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateComplaints(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"message":"invalid id",
		})
		return
	}

	var body struct {
		Title string		
		Description string	
		ImagePath string
	}


	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"failed to read body",
		})	
		return
	}

	var complaint models.Complaint


	if err := initializer.DB.First(&complaint, id).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "failed to get the complaint",
		})
		return
	}

	if complaint.CreatedBy != user.ID || user.Role != models.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"message":"not allowed to update this complaint",
		})
		return
	}

	complaint.Title = body.Title
	complaint.Description = body.Description
	complaint.ImagePath = body.ImagePath

	
	result := initializer.DB.Save(&complaint)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"fail to update the complaint",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}