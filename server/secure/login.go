package secure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/entity"
)

func BasicAuthLogin(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err == nil {
		fmt.Println(user)
		b, tmp := isValidUser(user)
		if b {
			json.NewEncoder(w).Encode(tmp)
			return
		}
	}

	fmt.Println("not ok")
	http.Error(w, "Unauthorized", http.StatusUnauthorized)

}
