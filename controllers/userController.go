package controllers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Balaji01-4D/ecoware-go/dto"
	"github.com/Balaji01-4D/ecoware-go/initializer"
	"github.com/Balaji01-4D/ecoware-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	var user dto.UserRegisterDto

	if err := c.BindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"cannot hash the password",
		})
		return
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hash),
		Role:     user.Role,
	}

	result := initializer.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"fail to create the user",
		})
		return
	}
	userResponse := dto.UserResponseDto{
		Name:  newUser.Name,
		Email: newUser.Email,
	}
	c.JSON(200, gin.H{
		"user": userResponse,
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

func getById(id int) *models.User {
	var user models.User
	result := initializer.DB.First(&user, id)

	if result.Error != nil {
		return nil
	}
	return &user
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
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)

}


func Login(c *gin.Context) {

	var body struct {
		Email string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"failed to read body",
		})
		return
	}
	var user models.User
	initializer.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"invalid email or password or user not found",
		})
		return
	
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"failed to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":tokenString,
	})


}