package ssh

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
	"maxiputz.github/sshManager/db/entity"
)

type SessionInfo struct {
	Client *ssh.Client
	User   string
	IP     string
}

func Connect(info entity.SSH) (SessionInfo, error) {
	conf := &ssh.ClientConfig{
		Timeout: time.Second * 2,
		User:    info.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(info.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", info.IPAddress+":22", conf)

	if err != nil {
		fmt.Println("issue with ssh dial ")
		return SessionInfo{}, err
	}

	return SessionInfo{
		Client: client,
		User:   info.User,
		IP:     info.IPAddress,
	}, err
}
