package entity

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
}

func FromInterfaceToUser(data map[string]interface{}) (User, error) {
	e1, ok1 := data["Name"].(string)
	e2, ok2 := data["Password"].(string)
	e3, ok3 := data["Email"].(string)

	if !ok1 || !ok2 || !ok3 {
		return User{}, errors.New("fail to parse User")
	}

	return User{
		Name:     e1,
		Password: e2,
		Email:    e3,
	}, nil
}
