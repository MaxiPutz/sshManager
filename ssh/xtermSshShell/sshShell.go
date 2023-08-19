package xtermsshshell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"maxiputz.github/sshManager/db/entity"
)

type sshShellController struct {
	websocketConn         *websocket.Conn
	sessionInputCannelMap map[string]sessionCHannel
}

var controller *sshShellController
var lock = &sync.Mutex{}

func Get() *sshShellController {
	if controller == nil {
		lock.Lock()
		defer lock.Unlock()
		if controller == nil {
			controller = &sshShellController{
				sessionInputCannelMap: make(map[string]sessionCHannel),
			}
			return controller
		}
		return controller
	}
	return controller
}

func (s *sshShellController) SetWebSocket(conn *websocket.Conn) {
	s.websocketConn = conn
}

func (s *sshShellController) AddSessionChannel(sc sessionCHannel) {
	s.sessionInputCannelMap[sc.uuid] = sc
}

func (s *sshShellController) WriteToSHHConnection(msg string, uuid string) {
	sc := s.sessionInputCannelMap[uuid]

	fmt.Printf("\"write to sshconnection \": %v\n", "write to sshconnection ")
	for k := range s.sessionInputCannelMap {
		fmt.Printf("sessionInputCannelMap %v\n", k)
	}

	sc.WriteInput(msg)
}

func (s *sshShellController) NewSession(info entity.SSH) (*ssh.Session, error) {
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
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sshShellController) OnMessageSocket(killChan chan bool) {
	socket := s.websocketConn
	go func() {
		for {
			_, msg, err := socket.ReadMessage()
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}

			if len(msg) == 0 {
				killChan <- true
				break
			}

			// csv := strings.Split(string(msg), ",")

			var sw writeStruct
			json.Unmarshal(msg, &sw)

			uuid := sw.Id
			command := sw.Msg
			fmt.Printf("uuid: %v\n", uuid)
			fmt.Printf("command: %v\n", command)

			Get().WriteToSHHConnection(command, uuid)

		}
	}()
	<-killChan
	fmt.Printf("\"kill chan true -> onmessage close\": %v\n", "kill chan true -> onmessage close")
}

func (s *sshShellController) SetPipe(session *ssh.Session, sessionID string, sc sessionCHannel) {
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to create stdout pipe: %s", err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatalf("Failed to create stderr pipe: %s", err)
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to create stdin pipe: %s", err)
	}

	// For xterm mode
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	term := "xterm-256color"

	session.RequestPty(term, 30, 40, modes)

	if err := session.Shell(); err != nil {
		log.Fatalf("Failed to start shell: %s", err)
	}

	write := socketWrite{
		channel: sessionID,
		socket:  Get().websocketConn,
	}

	go func() {
		io.Copy(write, stdout)
	}()
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	go func() {
		for {
			fmt.Println("waiting for input")
			msg := <-sc.GetChannelInput()
			fmt.Printf("channel input msg: %v\n", msg)
			io.Copy(stdin, bytes.NewReader([]byte(msg)))
		}
	}()

	<-sc.GetChannelKill()
}
