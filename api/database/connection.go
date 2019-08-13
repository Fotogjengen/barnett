package database

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

const (
	DBHOST = "localhost"
	DBPORT = 5432
	DBUSER = "postgres"
	DBPASS = ""
	DBNAME = "barnett_db"
)

var DB *sql.DB // Global variable

func Connect() {
	/*
	Returns sql.Db pointer
	Connects to database defined in constants
	*/
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", DBHOST, DBPORT, DBUSER, DBNAME)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	DB = db // Sets global variable DB to the db we connected to
}
