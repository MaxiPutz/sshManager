package secure

import (
	"fmt"
	"net/http"

	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/fn"
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
		if !isValidUser(username, password) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func isValidUser(username string, password string) bool {
	users := fn.Filter[entity.User](loadUsers(), func(b entity.User) bool {
		fmt.Println(b.Password)
		fmt.Println(b.Name)
		return b.Password == password && b.Name == username
	})

	return len(users) > 0
}
