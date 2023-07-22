package secure

import "maxiputz.github/sshManager/db/entity"

var (
	users []entity.User
)

func loadUsers() []entity.User {
	admin := entity.User{
		Name:     "admin",
		Password: "admin",
		Email:    "admin@admin.com",
	}

	return append(users, admin)
}
