package model

import "gorm.io/gorm"

type Voter struct {
	gorm.Model
	Contest Contest `gorm:"foreignKey:ID"`
	Code    string  `gorm:"unique;not null"`
	Used    bool    `gorm:"not null;default:false"`
}
