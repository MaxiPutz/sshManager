package db_info

import "fmt"

type DbInfo struct {
	Host     string
	UserName string
	Password string
	Dbname   string
	Port     int
	SSLMODE  string
}

func GetString() string {

	dbInfo := DbInfo{
		Host:     "localhost",
		UserName: "postgres",
		Password: "postgres",
		Dbname:   "sshmanager",
		Port:     5432,
		SSLMODE:  "disable",
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbInfo.Host, dbInfo.UserName, dbInfo.Password, dbInfo.Dbname, dbInfo.Port, dbInfo.SSLMODE)
}
