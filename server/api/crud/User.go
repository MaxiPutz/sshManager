package crud

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbcontainer "maxiputz.github/sshManager/db/db_container"
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
func UserDelete(w http.ResponseWriter, r *http.Request) {
	_db, err := db_singelton.GetDB()
	fmt.Println("in user delete request")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "cannot get db", http.StatusBadRequest)
		return
	}

	db := dbcontainer.DBContainer{
		G: _db,
	}

	user, ok := r.Context().Value("user").(entity.User)
	fmt.Printf("user: %v\n", user)
	if !ok {
		http.Error(w, "cannot get User", http.StatusBadRequest)
		return
	}

	// Delete action flows for the user
	flows, err := db.DeleteActionFlows(user)
	if err != nil {
		http.Error(w, "failed to delete action flows", http.StatusInternalServerError)
		return
	}
	for _, v := range flows {
		fmt.Printf("v.ActionName: %v\n", v.ActionName)
	}

	// Get and delete user SSH records
	u_s, err := db.GetUser_SSH(user)
	if err != nil {
		http.Error(w, "failed to get user SSH records", http.StatusInternalServerError)
		return
	}
	sshs, err := db.DeleteSSH(u_s)
	if err != nil {
		http.Error(w, "failed to delete user SSH records", http.StatusInternalServerError)
		return
	}
	for _, v := range sshs {
		fmt.Printf("v.IPAddress: %v\n", v.IPAddress)
	}

	// Delete user SSH mappings
	du_s, err := db.DeleteUser_SSH(user)
	if err != nil {
		http.Error(w, "failed to delete user SSH mappings", http.StatusInternalServerError)
		return
	}
	for _, v := range du_s {
		fmt.Printf("v.ssh: %v\n", v.SSH_id)
	}

	// Delete the user
	u, err := db.DeleteUser(user)
	if err != nil {
		http.Error(w, "failed to delete the user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)
}
