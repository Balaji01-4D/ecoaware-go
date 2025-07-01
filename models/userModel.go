package models

type UserRole string

const (
	RoleUser UserRole = "user"
	RoleAdmin UserRole = "admin"
)
type User struct {
	ID int 			`gorm:"primaryKey;autoIncrement" json:"id"`
	Name string		`gorm:"size:100" json:"name"`
	Email string	`gorm:"uniqueIndex" json:"email"`
	Password string	`gorm:"notnull" json:"password"`
	Role UserRole	`gorm:"notnull;type:user_role" json:"role"`
}

func (u *User) GetName() string {
	return  u.Name
}