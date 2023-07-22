package main

import (
	"fmt"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/server"
)

func main() {

	db, err := db_singelton.GetDB()
	if err != nil {
		fmt.Println("issue in to get the singelton db")
		panic(err)
	}
	defer db.Commit()
	var user entity.User

	db.First(&user, 1)

	fmt.Println(user.Name)

	server.InitServer()

}
