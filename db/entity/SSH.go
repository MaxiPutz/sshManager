package entity

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ParseError struct {
}

type SSH struct {
	gorm.Model
	IPAddress string `json:"IPAddress"`
	User      string `json:"User"`
	Password  string `json:"Password"`
	Key       string `json:"Key"`
}

func FromInterfaceToSSH(data map[string]interface{}) (SSH, error) {

	e1, ok1 := data["IPAddress"].(string)
	e2, ok2 := data["User"].(string)
	e3, ok3 := data["Password"].(string)
	e4, ok4 := data["Key"].(string)

	if !(ok1 && ok2 && ok3 && ok4) {
		fmt.Printf("data: %v\n", data)
		fmt.Printf("ok1: %v, ", ok1)
		fmt.Printf("ok2: %v, ", ok2)
		fmt.Printf("ok3: %v, ", ok3)
		fmt.Printf("ok4: %v,\n", ok4)

		return SSH{}, errors.New("Fail to parse")
	}

	return SSH{
		IPAddress: e1,
		User:      e2,
		Password:  e3,
		Key:       e4,
	}, nil
}
