package db

import (
	"fmt"
	"sysagent/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBCOnfig) string {

	cnnString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)

	if !cnf.EnableSSLMode {
		cnnString += " sslmode=disable"
	}

	return cnnString
	//return "user=postgres password=63316526 host=localhost port=5432 dbname=sysagent sslmode=disable"
}

func NewConnection(cnf *config.DBCOnfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbCon, nil
}
