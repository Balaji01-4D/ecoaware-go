package dto

import "github.com/Balaji01-4D/ecoware-go/models"


type UserRegisterDto struct {
	Name string 			`json:"name"`
	Email string 			`json:"email"`
	Password string 		`json:"password"`
	Role models.UserRole	`json:"role"`
}