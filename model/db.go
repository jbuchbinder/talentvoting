package model

import "gorm.io/gorm"

var DB *gorm.DB

func Migrate() error {
	var err error
	err = DB.AutoMigrate(&Contest{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Voter{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Option{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Vote{})
	if err != nil {
		return err
	}
	return nil
}
