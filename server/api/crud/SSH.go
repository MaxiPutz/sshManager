package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/ssh"
)

func ConnectionCheckHandler(w http.ResponseWriter, r *http.Request) {

	var sshEntity entity.SSH
	err := json.NewDecoder(r.Body).Decode(&sshEntity)
	if err != nil {
		fmt.Println("Failed to parse request body")
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	conn, err := ssh.Connect(sshEntity)
	if err != nil {
		fmt.Println("Faild to connect1")
		http.Error(w, "Faild to connect", http.StatusBadRequest)
		return
	}
	defer conn.Client.Close()

	str, err := ssh.Execute(conn, "ls")
	if err != nil {
		fmt.Println("Faild to connect2")
		http.Error(w, "Faild to connect", http.StatusBadRequest)
		return
	}
	fmt.Printf("str: %v\n", str)
	json.NewEncoder(w).Encode(str)
}

func CreateSSHHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into an SSH struct(w http.ResponseWriter, r *http.Request)
	var ssh entity.SSH
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&ssh)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Create the SSH record in the database

	db, err := db_singelton.GetDB()

	if err != nil {
		http.Error(w, "faild to get the db", http.StatusInternalServerError)
		return
	}

	result := db.Create(&ssh)
	if result.Error != nil {
		http.Error(w, "Failed to create SSH record", http.StatusInternalServerError)
		return
	}

	user, ok := r.Context().Value("user").(entity.User)
	if !ok {
		fmt.Println("not ok in the context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user_ssh := entity.User_SSH{
		User_id: (user.ID),
		SSH_id:  (ssh.ID),
	}

	result = db.Create(&user_ssh)
	if result.Error != nil {
		http.Error(w, "Failed to create user_ssh record", http.StatusInternalServerError)
		return
	}

	// Return the created SSH record in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ssh)
}

func GetSSHHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	// idStr := r.URL.Query().Get("id")

	user, ok := r.Context().Value("user").(entity.User)
	if !ok {
		fmt.Println("not ok in the context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}

	var user_sshs []entity.User_SSH
	result := db.Where("user_id = ?", user.ID).Find(&user_sshs)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "SSH record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve SSH record", http.StatusInternalServerError)
		}
		return
	}

	sshs := []entity.SSH{}
	for _, v := range user_sshs {
		sshID := v.SSH_id

		var ele entity.SSH
		db.Where("ID = ?", sshID).First(&ele)

		fmt.Printf("v: %v\n", v)
		fmt.Printf("sshID: %v\n", sshID)
		sshs = append(sshs, ele)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sshs)
}

func GetSSHHandlerByIP(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Query().Get("ip")

	// Get the database connection (db is assumed to be a global variable for the database connection)
	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}

	// Query the database to retrieve the SSH record with the given ID
	var ssh []entity.SSH
	result := db.First(&ssh, "ip_address = ?", idStr)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "SSH record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve SSH record", http.StatusInternalServerError)
		}
		return
	}

	// Return the SSH record in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ssh)
}

func UpdateSSHHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get the database connection (db is assumed to be a global variable for the database connection)
	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}

	// Check if the SSH record with the given ID exists
	var ssh entity.SSH
	result := db.First(&ssh, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "SSH record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve SSH record", http.StatusInternalServerError)
		}
		return
	}

	// Parse the request body into an SSH struct
	var updatedSSH entity.SSH
	err = json.NewDecoder(r.Body).Decode(&updatedSSH)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Update the existing SSH record with the new values

	if ssh.IPAddress != "" {
		ssh.IPAddress = updatedSSH.IPAddress
	}

	if ssh.User != "" {
		ssh.User = updatedSSH.User
	}

	if ssh.Password != "" {
		ssh.Password = updatedSSH.Password
	}

	if ssh.Key != "" {
		ssh.Key = updatedSSH.Key
	}

	result = db.Save(&ssh)
	if result.Error != nil {
		http.Error(w, "Failed to update SSH record", http.StatusInternalServerError)
		return
	}

	// Return the updated SSH record in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ssh)

}

func DeleteSSHHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get the database connection (db is assumed to be a global variable for the database connection)
	db, err := db_singelton.GetDB()
	if err != nil {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}

	// Check if the SSH record with the given ID exists
	var ssh entity.SSH
	result := db.First(&ssh, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "SSH record not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve SSH record", http.StatusInternalServerError)
		}
		return
	}

	// Delete the SSH record from the database
	result = db.Delete(&ssh, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete SSH record", http.StatusInternalServerError)
		return
	}

	// Return a success message in the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SSH record deleted successfully"))
}
