package main

import (
	"fmt"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/server"
	"maxiputz.github/sshManager/ssh"
	testdata "maxiputz.github/sshManager/testData"
)

func main() {

	db, err := db_singelton.GetDB()
	if err != nil {
		fmt.Println("issue in to get the singelton db")
		panic(err)
	}
	defer db.Commit()

	//TestSSH()
	server.InitServer()
}

func TestSSH() {
	conn, err := ssh.Connect(testdata.Info())
	if err != nil {
		panic(err)
	}
	defer conn.Client.Close()

	out, err := ssh.Execute(conn, "pwd")
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %v\n", out)

	out, err = ssh.Execute(conn, "ls /Users/max/workspace/react/client-side-encryption-app")
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %v\n", out)

	// err = ssh.CopyDirToHost(conn, "/Users/max/workspace/react/client-side-encryption-app/src", "/Users/max/workspace/sshManagerFrontend/src")
	err = ssh.CopyDirFromRemote(conn, "/Users/max/workspace/react/client-side-encryption-app/public", "/Users/max/workspace/sshManagerFrontend/public")
	if err != nil {
		panic(err)
	}

	fmt.Printf("out: %v\n", out)

}
