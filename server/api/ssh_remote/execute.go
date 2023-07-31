package sshremote

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/ssh"
)

func Execute(w http.ResponseWriter, r *http.Request) {

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("Faild to get the interface")
		http.Error(w, "Faild to get the interface", http.StatusBadRequest)
		return
	}
	sshInfo, err := entity.FromInterfaceToSSH(data)

	if err != nil {
		fmt.Println("Faild to get sshinfo")
		http.Error(w, "Faild to get sshinfo", http.StatusBadRequest)
		return
	}

	command := entity.SSHCommand{}
	tmp, ok := data["Command"].(string)
	if !ok {
		fmt.Println("Faild to get sshcommand")

		http.Error(w, "Faild to get sshcommand", http.StatusBadRequest)
		return
	}

	command.Command = tmp

	session, err := ssh.Connect(sshInfo)
	if err != nil {
		fmt.Println("Cannot connect")
		http.Error(w, "Cannot connect", http.StatusBadRequest)
		return
	}

	msg, err := ssh.Execute(session, command.Command)
	if err != nil {
		fmt.Println("Cannot exe")
		http.Error(w, "Cannot execute", http.StatusBadRequest)
		return
	}
	fmt.Printf("msg: %v\n", msg)
	json.NewEncoder(w).Encode(msg)

}
