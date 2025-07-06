package models

import "time"

type RefreshSession struct {
	ID	uint			`gorm:"primaryKey"`
	UserId uint			`gorm:"index"`
	Token	string		`gorm:"uniqueIndex"`
	ExpiresAt time.Time
	CreatedAt time.Time
}