package controllers

import (
	"net/http"
	"strconv"
	"time"
	"fmt"

	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
)

func GetAllComplaints(c *gin.Context) {

	
	var complaints []models.Complaint

	result := initializer.DB.Find(&complaints)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to retirve",
		})
		return
	}

	c.JSON(http.StatusOK, complaints)
}

func AddComplaints(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	title := c.PostForm("title")
	description := c.PostForm("description")
	categoryIDStr := c.PostForm("categoryId") 
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid category ID"})
		return
	}

	file, err := c.FormFile("image")
	var imagePath string
	if err == nil {
		imagePath = fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)
		err = c.SaveUploadedFile(file, imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save file"})
			return
		}
	}

	complaint := models.Complaint{
		Title:      title,
		Description: description,
		ImagePath:   imagePath,
		CreatedBy:   user.ID,
		CreatedAt:   time.Now(),
		CategoryID:  uint(categoryID),
	}

	if result := initializer.DB.Create(&complaint); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to add complaint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint added successfully", "complaint": complaint})
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

func GetMyComplaints(c *gin.Context) {

	user := c.MustGet("user").(models.User)


	var complaints []models.Complaint

	err := initializer.DB.
		Preload("User").
		Preload("Category").
		Where("created_by", user.ID).
		Find(&complaints).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to fetech the complaints",
		})
	}

	c.JSON(http.StatusOK,complaints)
}

func GetComplaintByID(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"message":"invalid id",
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
			"message":"not allowed to retrive this complaint",
		})
		return
	}

	c.JSON(http.StatusOK, complaint)

}


func UpdateComplaintStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(400, gin.H{
			"message":"invalid id",
		})
		return
	}

	var body struct {
		Status models.Status `json:"status" binding:"required"`
	}		

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"cannot read the body",
		})
		return
	}

	if err := initializer.DB.Model(&models.Complaint{}).Where("id = ?", id).Update("status", body.Status).Error;
		err != nil {
			c.JSON(http.StatusNotFound, gin.H{
					"message":"failed to update the status",
				})
		return
	}
			
	c.JSON(http.StatusOK, gin.H{
		"message":"status updated successfully",
	})
	
}