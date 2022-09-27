package model

import "gorm.io/gorm"

type Contest struct {
	gorm.Model
	Slug        string
	Name        string
	Description string
}
