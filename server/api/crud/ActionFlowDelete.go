package crud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
)

func ActionFlowDelete(w http.ResponseWriter, r *http.Request) {
	type resType struct {
		UUID string `json:"UUID"`
	}
	uuid := resType{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &uuid)
	fmt.Printf("uuid: %v\n", uuid)
	if err != nil {
		http.Error(w, "cannot read unmarshal body", http.StatusBadRequest)
		return
	}

	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "cannot open db", http.StatusBadRequest)
		return
	}

	af := []entity.ActionFlow{}
	db.Where(&entity.ActionFlow{UUID: uuid.UUID}).Find(&af)

	for _, v := range af {
		fmt.Printf("v: %v\n", v)
		result := db.Delete(&entity.ActionFlow{}, v.ID)
		if result.Error != nil {
			fmt.Println("cannnot delete the entriy")
			fmt.Printf("result.Error: %v\n", result.Error)
			http.Error(w, "cannnot delete the entriy", http.StatusBadRequest)
			return
		}
	}
	json.NewEncoder(w).Encode(af)
}
