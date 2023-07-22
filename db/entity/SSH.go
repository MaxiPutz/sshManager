package entity

import "gorm.io/gorm"

type SSH struct {
	gorm.Model
	IPAddress string
	User      string
	Password  string
	Key       string
}
