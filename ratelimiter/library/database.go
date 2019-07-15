package library

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 15432
	user     = "local"
	password = "local@007"
	dbname   = "ratelimiter"
)

var db *sql.DB

func SetupDSN() string {
	var psqlinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlinfo
}

func SetupConnection() *sql.DB {

	var err error
	db, err = sql.Open("postgres", SetupDSN())

	if err != nil {
		log.Panic(err)
	}

	return db

}

func CloseConnection() {
	db.Close()
}

func PingDb() bool {
	err := db.Ping()
	if err != nil {
		log.Panic(err)
	}

	return true
}

func GetDBConnection() *sql.DB {
	return db
}
