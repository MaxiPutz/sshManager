package server

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"maxiputz.github/sshManager/server/api"
	"maxiputz.github/sshManager/server/secure"
)

func InitServer() {
	mx := http.NewServeMux()

	fs := http.FileServer(http.Dir("./server/public"))
	mx.Handle("/", fs)

	mx.HandleFunc("/ssh/create", secure.BasicAuth(api.CreateSSHHandler))
	mx.HandleFunc("/ssh/get", secure.BasicAuth(api.GetSSHHandler))
	mx.HandleFunc("/ssh/getIp", secure.BasicAuth(api.GetSSHHandlerByIP))
	mx.HandleFunc("/ssh/update/{id}", secure.BasicAuth(api.UpdateSSHHandler))
	mx.HandleFunc("/ssh/delete/{id}", secure.BasicAuth(api.DeleteSSHHandler))
	mx.HandleFunc("/user/create", api.UserCreate)

	log.Println("Server started on :8080")

	handler := cors.AllowAll().Handler(mx)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
