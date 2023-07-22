package server

import (
	"log"
	"net/http"

	api "maxiputz.github/sshManager/server/API"
	"maxiputz.github/sshManager/server/secure"
)

func InitServer() {

	fs := http.FileServer(http.Dir("./server/public"))
	http.Handle("/", fs)

	http.HandleFunc("/ssh/create", secure.BasicAuth(api.CreateSSHHandler))
	http.HandleFunc("/ssh/get", secure.BasicAuth(api.GetSSHHandler))
	http.HandleFunc("/ssh/getIp", secure.BasicAuth(api.GetSSHHandlerByIP))
	http.HandleFunc("/ssh/update/{id}", secure.BasicAuth(api.UpdateSSHHandler))
	http.HandleFunc("/ssh/delete/{id}", secure.BasicAuth(api.DeleteSSHHandler))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
