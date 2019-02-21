package database

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	DBHOST = "db"
	DBUSER = "postgres-dev"
	DBPASS = "qwer1234"
	DBNAME = "dev"
)

var DB *sql.DB // Global variable

func Connect() {
	/*
	Returns sql.Db pointer
	Connects to database defined in constants
	*/
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DBHOST, DBUSER, DBPASS, DBNAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	DB = db // Sets global variable DB to the db we connected to
}
