package db_singelton

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"maxiputz.github/sshManager/db/db_info"
	"maxiputz.github/sshManager/db/entity"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() (*gorm.DB, error) {
	var err error

	dbOnce.Do(func() {
		fmt.Println("try to create the db")
		dsn := db_info.GetString()

		fmt.Printf("dsn: %v\n", dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("error in db string")
		}

		db.AutoMigrate(&entity.User{})
		db.AutoMigrate(&entity.SSH{})

	})

	return db, err
}
