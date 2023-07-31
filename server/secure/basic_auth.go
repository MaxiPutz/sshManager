package secure

import (
	"context"
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/db_singelton"
	"maxiputz.github/sshManager/db/entity"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has the "Authorization" header

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("no basic auth header")
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the username and password from the "Authorization" header
		username, password, ok := r.BasicAuth()
		fmt.Println("basic auth header ist ", username, password)

		if !ok {
			fmt.Println("not ok")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the provided credentials are valid
		b, user := isValidUser(entity.User{
			Name:     username,
			Password: password,
		})
		if !b {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next(w, r.WithContext(ctx))
	}
}

func isValidUser(user entity.User) (bool, entity.User) {
	db, err := db_singelton.GetDB()
	if err != nil {
		return false, entity.User{}
	}

	var resUser entity.User
	res := db.Where("name = ? ", user.Name).First(&resUser)

	if res.Error != nil {
		fmt.Println("no user found")
		return false, entity.User{}
	}

	b := CompareHashStrWithStr(resUser.Password, user.Password)

	if b == false {
		fmt.Println("hash doesnt fit")
		return false, entity.User{}
	}

	return true, resUser
}
