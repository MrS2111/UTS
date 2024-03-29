package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connect() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	fmt.Println(dbHost)
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/uts_pbp")
	if err != nil {
		log.Fatal(err)
	}
	return db
}