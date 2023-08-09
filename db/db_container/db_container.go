package dbcontainer

import (
	"fmt"

	"gorm.io/gorm"
	"maxiputz.github/sshManager/db/entity"
	"maxiputz.github/sshManager/fn"
)

type DBContainer struct {
	G *gorm.DB
}

func (db *DBContainer) DeleteActionFlows(u entity.User) ([]entity.ActionFlow, error) {
	g := db.G

	actionFLows := []entity.ActionFlow{}

	err := g.Where(&entity.ActionFlow{User_Id: u.ID}).Delete(&actionFLows)

	return actionFLows, err.Error
}

func (db *DBContainer) GetUser_SSH(u entity.User) ([]entity.User_SSH, error) {
	g := db.G

	user_ssh := []entity.User_SSH{}

	err := g.Where(&entity.User_SSH{User_id: u.ID}).Find(&user_ssh)

	return user_ssh, err.Error
}

func (db *DBContainer) DeleteSSH(u_s []entity.User_SSH) ([]entity.SSH, error) {
	g := db.G

	ssh := [](entity.SSH){}
	errs := []error{}

	for _, v := range u_s {
		tmp := entity.SSH{}
		err := g.First(&tmp, v.SSH_id)
		ssh = append(ssh, tmp)
		errs = append(errs, err.Error)
		fmt.Printf("tmp: %v\n", tmp)
		fmt.Printf("v: %v\n", v)
	}

	errs = fn.Filter[error](errs, func(ele error) bool {
		return ele != nil
	})
	if len(errs) != 0 {
		return ssh, errs[0]
	}

	err := g.Delete(ssh)
	return ssh, err.Error
}

func (db *DBContainer) DeleteUser_SSH(u entity.User) ([]entity.User_SSH, error) {
	g := db.G

	user_ssh := []entity.User_SSH{}

	err := g.Where(&entity.User_SSH{User_id: u.ID}).Delete(&user_ssh)

	return user_ssh, err.Error
}

func (db *DBContainer) DeleteUser(u entity.User) (entity.User, error) {

	g := db.G
	err := g.Delete(&u)

	return u, err.Error
}
