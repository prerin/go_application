package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
)

type Person struct {
	Id string 
	Name string
}

func InsertValue(name string) {
	db, err := sql.Open("sqlite3", "./first.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO manager (name) values(?)`,
	name)
	if err != nil {
		log.Fatal(err)
	}	
}

func GetAllPerson() []Person {
	db, err := sql.Open("sqlite3", "./first.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT * FROM manager`,
	)
	if err != nil {
		log.Fatal(err)
	}	
	defer rows.Close()

	person_ary := []Person{}
	for rows.Next() {
		var id string
		var name string

		if err := rows.Scan(&id, & name); err != nil {
			log.Fatal("rows.Scan()", err)
			return []Person{}
		}
		person_ary = append(person_ary, Person{id, name})
	}
	return person_ary
}

func DeletePerson(id string) {
	db, err := sql.Open("sqlite3", "./first.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM manager WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
}