package models

import "time"


type Complaint struct {
	ID uint 			`gorm:"primaryKey" json:"id"`
	Title string		
	Description string	
	CreatedAt time.Time	
	CreatedBy uint
	User User 			`gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}