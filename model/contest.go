package model

import "gorm.io/gorm"

type Contest struct {
	gorm.Model
	Code        string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	MaxVotes    int    `gorm:"default:5"`
}
