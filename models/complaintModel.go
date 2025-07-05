package models

import "time"

type Complaint struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      Status    `gorm:"type:status;default:'PENDING';not null" json:"status"` 
	ImagePath   string    `gorm:"size:255" json:"imagePath"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   uint      `json:"createdBy"`
	User        User      `gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
	CategoryID  uint      `json:"categoryId"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
}
