package controller

import (
	"database/sql"
	"log"
)

func openConnection() *sql.DB {
	// open DB connection
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
