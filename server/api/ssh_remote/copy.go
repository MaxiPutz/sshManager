package sshremote

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/ssh"
)

func CopyFileFromRemote(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err1 := json.NewDecoder(r.Body).Decode(&data)

	action, err2 := entity.FromInterfaceToSSHCopy(data)
	_info, err3 := entity.FromInterfaceToSSH(data)

	if err1 != nil || err2 != nil || err3 != nil {
		http.Error(w, "fail to parse json", http.StatusBadRequest)
		return
	}

	fmt.Printf("action: %v\n", action)
	fmt.Printf("_info: %v\n", _info)
	info, err := ssh.Connect(_info)
	if err != nil {
		http.Error(w, "fail to connect to ssh server", http.StatusBadRequest)
		return
	}

	ssh.CopyFileFromRemote(info, action.Source_Dir, action.Destination_Dir)

	json.NewEncoder(w).Encode("copy was success")
}

func CopyFileToRemote(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err1 := json.NewDecoder(r.Body).Decode(&data)

	action, err2 := entity.FromInterfaceToSSHCopy(data)
	_info, err3 := entity.FromInterfaceToSSH(data)

	if err1 != nil || err2 != nil || err3 != nil {
		http.Error(w, "fail to parse json", http.StatusBadRequest)
		return
	}

	fmt.Printf("action: %v\n", action)
	fmt.Printf("_info: %v\n", _info)
	info, err := ssh.Connect(_info)
	if err != nil {
		http.Error(w, "fail to connect to ssh server", http.StatusBadRequest)
		return
	}

	ssh.CopyFileToRemote(info, action.Source_Dir, action.Destination_Dir)

	json.NewEncoder(w).Encode("copy was success")
}
