package model

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	Contest Contest `gorm:"foreignKey:ID"`
	Voter   Voter   `gorm:"foreignKey:ID"`
	Option  Option  `gorm:"foreignKey:ID"`
}
