package apishell

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"maxiputz.github/sshManager/db/entity"
	sshshell "maxiputz.github/sshManager/ssh/xtermSshShell"
)

func HandleNewSSHShellConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\"helloworld\": %v\n", "helloworld")

	type UUIDStruct struct {
		UUID string `json:"UUID"`
	}
	var uuid UUIDStruct
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, &uuid)
	fmt.Printf("uuid: %v\n", uuid)

	var sshInfo entity.SSH
	json.Unmarshal(b, &sshInfo)

	fmt.Printf("sshInfo: %v\n", sshInfo)

	session, err := sshshell.Get().NewSession(sshInfo)
	if err != nil {
		http.Error(w, "not able to make a session", http.StatusBadRequest)
		return
	}

	sc := sshshell.NewSessionChannel(uuid.UUID)

	go sshshell.Get().SetPipe(session, uuid.UUID, sc)
	sshshell.Get().AddSessionChannel(sc)
	json.NewEncoder(w).Encode("hello wolrd")
}

func StreamConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stream conn")
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, _ := upgrader.Upgrade(w, r, nil)
	fmt.Printf("conn.RemoteAddr(): %v\n", conn.RemoteAddr())

	sshshell.Get().SetWebSocket(conn)
	killChan := make(chan bool)
	go sshshell.Get().OnMessageSocket(killChan)

}
