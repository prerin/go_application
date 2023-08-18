package main

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"log"
)

func create_table() {
	db, err := sql.Open("sqlite3", "./first.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tableName := "manager"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL)`, tableName)
	
	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}	
}