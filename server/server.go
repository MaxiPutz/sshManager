package server

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	apishell "maxiputz.github/sshManager/server/api/apiShell"
	"maxiputz.github/sshManager/server/api/crud"
	sshremote "maxiputz.github/sshManager/server/api/ssh_remote"
	"maxiputz.github/sshManager/server/secure"
)

func InitServer() {
	mx := http.NewServeMux()

	fs := http.FileServer(http.Dir("./server/public/build"))
	mx.Handle("/", fs)

	mx.HandleFunc("/ssh/create", secure.BasicAuth(crud.CreateSSHHandler))
	mx.HandleFunc("/ssh/checkConnection", secure.BasicAuth(crud.ConnectionCheckHandler))
	mx.HandleFunc("/ssh/get", secure.BasicAuth(crud.GetSSHHandler))
	mx.HandleFunc("/ssh/getIp", secure.BasicAuth(crud.GetSSHHandlerByIP))
	mx.HandleFunc("/ssh/update/{id}", secure.BasicAuth(crud.UpdateSSHHandler))
	mx.HandleFunc("/ssh/delete", secure.BasicAuth(crud.DeleteSSHHandler))

	mx.HandleFunc("/ssh/exe", secure.BasicAuth((sshremote.Execute)))

	mx.HandleFunc("/ssh/copyFileFromRemote", secure.BasicAuth((sshremote.CopyFileFromRemote)))
	mx.HandleFunc("/ssh/copyFileToRemote", secure.BasicAuth((sshremote.CopyFileToRemote)))

	mx.HandleFunc("/ssh/actionFlow/create", secure.BasicAuth((crud.CreateActionFlowHandler)))
	mx.HandleFunc("/ssh/actionFlow/getAll", secure.BasicAuth((crud.ActionFlowGetAll)))
	mx.HandleFunc("/ssh/actionFlow/delete", secure.BasicAuth(crud.ActionFlowDelete))

	mx.HandleFunc("/sshShell/newConnection", apishell.HandleNewSSHShellConnection)
	mx.HandleFunc("/sshShell/socket", apishell.StreamConn)

	mx.HandleFunc("/user/create", crud.UserCreate)
	mx.HandleFunc("/user/delete", secure.BasicAuth(crud.UserDelete))

	mx.HandleFunc("/login", secure.BasicAuthLogin)

	log.Println("Server started on :8080")

	handler := cors.AllowAll().Handler(mx)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
