package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	apishell "maxiputz.github/sshManager/server/api/apiShell"
	"maxiputz.github/sshManager/server/api/crud"
	sshremote "maxiputz.github/sshManager/server/api/ssh_remote"
	"maxiputz.github/sshManager/server/secure"
)

func getBasePath() string {
	base := os.Getenv("BASE_PATH")

	// Default: leer (= root)
	if base == "" {
		return ""
	}

	// muss mit / anfangen
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}

	// trailing slash entfernen
	base = strings.TrimRight(base, "/")

	return base
}

func InitServer() {
	basePath := getBasePath()
	mx := http.NewServeMux()

	fs := http.StripPrefix(
		basePath,
		http.FileServer(http.Dir("./server/public/build")),
	)
	mx.Handle(basePath+"/", fs)

	mx.HandleFunc(basePath+"/ssh/create", secure.BasicAuth(crud.CreateSSHHandler))
	mx.HandleFunc(basePath+"/ssh/checkConnection", secure.BasicAuth(crud.ConnectionCheckHandler))
	mx.HandleFunc(basePath+"/ssh/get", secure.BasicAuth(crud.GetSSHHandler))
	mx.HandleFunc(basePath+"/ssh/getIp", secure.BasicAuth(crud.GetSSHHandlerByIP))
	mx.HandleFunc(basePath+"/ssh/update/{id}", secure.BasicAuth(crud.UpdateSSHHandler))
	mx.HandleFunc(basePath+"/ssh/delete", secure.BasicAuth(crud.DeleteSSHHandler))

	mx.HandleFunc(basePath+"/ssh/exe", secure.BasicAuth((sshremote.Execute)))

	mx.HandleFunc(basePath+"/ssh/copyFileFromRemote", secure.BasicAuth((sshremote.CopyFileFromRemote)))
	mx.HandleFunc(basePath+"/ssh/copyFileToRemote", secure.BasicAuth((sshremote.CopyFileToRemote)))

	mx.HandleFunc(basePath+"/ssh/actionFlow/create", secure.BasicAuth((crud.CreateActionFlowHandler)))
	mx.HandleFunc(basePath+"/ssh/actionFlow/getAll", secure.BasicAuth((crud.ActionFlowGetAll)))
	mx.HandleFunc(basePath+"/ssh/actionFlow/delete", secure.BasicAuth(crud.ActionFlowDelete))

	mx.HandleFunc(basePath+"/sshShell/newConnection", apishell.HandleNewSSHShellConnection)
	mx.HandleFunc(basePath+"/sshShell/socket", apishell.StreamConn)

	mx.HandleFunc(basePath+"/user/create", crud.UserCreate)
	mx.HandleFunc(basePath+"/user/delete", secure.BasicAuth(crud.UserDelete))

	mx.HandleFunc(basePath+"/login", secure.BasicAuthLogin)

	log.Println("Server started on :8081")

	handler := cors.AllowAll().Handler(mx)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
