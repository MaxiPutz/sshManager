package entity

import "gorm.io/gorm"

type User_SSH struct {
	gorm.Model
	User_id uint
	SSH_id  uint
}
