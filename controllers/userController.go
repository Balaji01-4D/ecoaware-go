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

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message":"i am logged in",
		"profile":user,
	})
}

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


func UpdateUserByUser(c *gin.Context) {
	loginedUser := c.MustGet("user").(models.User)

	var userRecord models.User

	if err := initializer.DB.First(&userRecord, loginedUser.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to get the user",
		})
		return
	}

	var body struct {
		Name string 		`json:"name"`
		Email string 		`json:"email"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"failed to read body",
		})
		return
	}

	if err := initializer.DB.Model(&loginedUser).
	Update("name", body.Name).
	Update("email", body.Email).
	Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"successfully updated",
	})
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

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{})


}

func Me(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))


	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok  && token.Valid{
		var user models.User

		initializer.DB.First(&user, "email = ?", claims["sub"])

		if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		})
	}else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
	
	}

}

func UpdatePassword(c *gin.Context) {
	loginedUser := c.MustGet("user").(models.User)

	var userRecord models.User

	if err := initializer.DB.First(&userRecord, loginedUser.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to get the user",
		})
		return
	}

	var body struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword string 		`json:"newPassword"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" :"failed to read body",
		})
		return
	}

	if userRecord.Password != body.CurrentPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"wrong password",
		})
		return
	}

	if err := initializer.DB.Model(&loginedUser).
	Update("password", body.NewPassword).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"failed to update the password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"successfully updated the password",
	})
}