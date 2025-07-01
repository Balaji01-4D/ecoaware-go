package models

type Category struct {
	ID int8 		`gorm:"primaryKey;autoIncrement" json:"id"`
	Category string	`gorm:"notnull" json:"category"`
}