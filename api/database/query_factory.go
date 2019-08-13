package database

import (
	"database/sql"
	"log"
)

func Query(q string) *sql.Rows {
	/*
	Return rows from query to DB
	Error handling is handled here in the factory
	*/
	// prepare query
	stmt, err := DB.Prepare(q)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	return rows
}

func QueryOne(q string) *sql.Row {
	stmt, err := DB.Prepare(q)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	return row
}

func Insert(q string) error {
	stmt, err := DB.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // TODO: handle error

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
