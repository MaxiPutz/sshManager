package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
)

func CreateActionFlowHandler(w http.ResponseWriter, r *http.Request) {
	var container map[string]interface{}

	json.NewDecoder(r.Body).Decode(&container)

	ActionName, ok := container["ActionName"].(string)
	if !ok {
		http.Error(w, "faild to open db", http.StatusBadRequest)
		return
	}
	fmt.Printf("ActionName: %v\n", ActionName)

	user, err1 := entity.FromInterfaceToUser(container)

	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "faild to open db", http.StatusBadRequest)
		return
	}

	res := db.Where("name = ?", user.Name).First(&user)

	fmt.Printf("res: %v\n", res)

	fmt.Printf("ssh: %v\n", user)

	commands, _ := container["commands"].([]interface{})

	sshCommand := []entity.SSHCommand{}
	sshCopy := []entity.SSHCopy{}
	for _, v := range commands {
		val, ok1 := v.(map[string]interface{})
		if ok1 {
			v1, ok2 := val["Action"].(string)
			if ok2 {
				if v1 == "SSHCommand" {
					v2, err := entity.FromInterfaceToSSHCommand(val)
					sshCommand = append(sshCommand, v2)
					if err != nil {
						err1 = err
					}
				} else {
					v2, err := entity.FromInterfaceToSSHCopy(val)
					sshCopy = append(sshCopy, v2)
					if err != nil {
						err1 = err
					}
				}
			} else {
				err1 = errors.New("commands faild to Parse")
			}

		} else {
			err1 = errors.New("commands faild to Parse")
		}

	}

	_sshInfos, ok1 := container["sshInfos"].([]interface{})
	sshInfos := []entity.SSH{}
	for _, v := range _sshInfos {
		val, ok1 := v.(map[string]interface{})

		if ok1 {
			val1, err := entity.FromInterfaceToSSH(val)

			_sshInfos = append(_sshInfos, val1)
			if err != nil {
				err1 = err
			}

		} else {
			err1 = errors.New("commands faild to Parse")
		}

	}

	fmt.Printf("sshInfos: %v\n", sshInfos)

	if err1 != nil || !ok1 {
		fmt.Printf("err1: %v\n", err1)
		http.Error(w, "faild to parse", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	actionFlow := []entity.ActionFlow{}
	for _, v := range sshCommand {
		actionFlow = append(actionFlow, entity.ActionFlow{
			User_Id:         user.ID,
			Command:         v.Command,
			Target:          v.Target,
			Action:          v.Action,
			Source_Dir:      "",
			Destination_Dir: "",
			ActionFlow_uuid: id,
			Index:           v.Index,
			ActionName:      ActionName,
		})
	}

	for _, v := range sshCopy {
		actionFlow = append(actionFlow, entity.ActionFlow{
			User_Id:         user.ID,
			Command:         "",
			Target:          v.Target,
			Action:          v.Action,
			Source_Dir:      v.Source_Dir,
			Destination_Dir: v.Destination_Dir,
			ActionFlow_uuid: id,
			Index:           v.Index,
			ActionName:      ActionName,
		})
	}

	actionFlowDBInsert(actionFlow)

	json.NewEncoder(w).Encode(actionFlow)
}

func actionFlowDBInsert(actionFlow []entity.ActionFlow) error {
	db, err := db_singelton.GetDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	for _, v := range actionFlow {
		result := db.Create(&v)
		fmt.Printf("result: %v\n", result)
	}

	return nil
}
