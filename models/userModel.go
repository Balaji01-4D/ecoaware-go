package models

type UserRole string

const (
	RoleUser UserRole = "user"
	RoleAdmin UserRole = "admin"
)


type User struct {
	ID         uint        		`gorm:"primaryKey" json:"id"`
	Name       string      		`gorm:"size:100;not null" json:"name"`
	Email      string      		`gorm:"uniqueIndex;not null" json:"email"`
	Password   string      		`gorm:"not null" json:"-"`
	Role       UserRole    		`gorm:"type:user_role;not null" json:"role"`
	Complaints []Complaint 		`gorm:"foreignKey:CreatedBy" json:"complaints"`
}

