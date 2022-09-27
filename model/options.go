package model

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Contest Contest `gorm:"foreignKey:ID"`
	Name    string
}
