package model

import "gorm.io/gorm"

type Voter struct {
	gorm.Model
	Code string
	Used bool
}
