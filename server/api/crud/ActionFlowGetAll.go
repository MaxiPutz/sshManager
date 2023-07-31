package crud

import (
	"encoding/json"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
)

func ActionFlowGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the get all request")
	user, ok := r.Context().Value("user").(entity.User)

	fmt.Printf("user: %v\n", user)

	db, err := db_singelton.GetDB()

	if !ok || err != nil {
		fmt.Printf("fail to get User \t ok: %v, ", ok)
		fmt.Printf("err: %v\n", err)
		http.Error(w, "fail to get user", http.StatusBadRequest)
		return
	}

	var actionflows []entity.ActionFlow
	result := db.Where("user_id = ?", user.ID).Find(&actionflows)

	if result.Error != nil {
		fmt.Printf("result: %v\n", result)
		http.Error(w, "fail to get user", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(actionflows)
}
