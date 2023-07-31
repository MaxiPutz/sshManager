package crud

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/server/secure"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello testing")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "db singelton issue", http.StatusBadRequest)
		return
	}

	var resUser entity.User
	res := db.Where("name = ?", user.Name).First(&resUser)

	if res.Error == nil {
		fmt.Println(resUser)
		http.Error(w, "user exist", http.StatusBadRequest)
		return
	}

	hash, err := secure.HashString(user.Password)
	if err != nil {
		http.Error(w, "Hashing failed", http.StatusBadRequest)
		return
	}
	user.Password = hash

	db.Create(&user)

	json.NewEncoder(w).Encode(user)
}
