package entity

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ActionFlow struct {
	gorm.Model
	ActionFlow_uuid string
	ActionName      string
	User_Id         uint
	Index           int
	Command         string
	Action          string
	Target          string
	Source_Dir      string
	Destination_Dir string
}

type SSHCommand struct {
	Command string
	Action  string
	Target  string
	Index   int
}

type SSHCopy struct {
	Action          string
	Target          string
	Source_Dir      string
	Destination_Dir string
	Index           int
}

func FromInterfaceToSSHCommand(data map[string]interface{}) (SSHCommand, error) {
	e1, ok1 := data["Command"].(string)
	e2, ok2 := data["Action"].(string)
	e3, ok3 := data["Target"].(string)
	e4, ok4 := data["Index"].(float64)

	fmt.Printf("data: %v\n", data)
	if !ok1 || !ok2 || !ok3 || !ok4 {
		fmt.Printf("ok1: %v, ", ok1)
		fmt.Printf("ok2: %v, ", ok2)
		fmt.Printf("ok3: %v, ", ok3)
		fmt.Printf("ok4: %v, ", ok4)
		return SSHCommand{}, errors.New("fail to parse sshcomand")
	}

	return SSHCommand{
		Command: e1,
		Action:  e2,
		Target:  e3,
		Index:   int(e4),
	}, nil
}

func TargetFile() string {
	return "File"
}
func TargetDirectory() string {
	return "Dir"
}

func FromInterfaceToSSHCopy(data map[string]interface{}) (SSHCopy, error) {

	e1, ok1 := data["Source_Dir"].(string)
	e2, ok2 := data["Destination_Dir"].(string)
	e3, ok3 := data["Action"].(string)
	e4, ok4 := data["Target"].(string)
	e5, ok5 := data["Index"].(float64)

	fmt.Printf("e5: %v\n", e5)
	fmt.Printf("data: %v\n", data)

	// index, err := strconv.Atoi(e5)

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
		fmt.Printf("ok1: %v, ", ok1)
		fmt.Printf("ok2: %v, ", ok2)
		fmt.Printf("ok3: %v, ", ok3)
		fmt.Printf("ok4: %v, ", ok4)
		fmt.Printf("ok5: %v, ", ok5)
		// fmt.Printf("err: %v\n", err)

		return SSHCopy{}, errors.New("fail to parse sshcopy")
	}

	return SSHCopy{
		Source_Dir:      e1,
		Destination_Dir: e2,
		Action:          e3,
		Target:          e4,
		Index:           int(e5),
	}, nil
}
