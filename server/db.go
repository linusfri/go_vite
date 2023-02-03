package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Client *sql.DB
}

type Person struct {
	Name   string
	Adress string
}

func NewDatabase() *DB {
	client, err := sql.Open("mysql", "root:ippenutt1@(db:3306)/db_data")
	if err != nil {
		panic(err)
	}
	return &DB{Client: client}
}

func (db *DB) init() (rows *sql.Rows, err error) {
	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id int AUTO_INCREMENT,
			name varchar(20),
			adress varchar(100),
			PRIMARY KEY (id)
		);
	`

	return db.Client.Query(sql)
}

func (db *DB) insert(person Person) error {
	_, err := db.Client.Query(`
		INSERT INTO users (name, adress)
		VALUES (?, ?)
	`, person.Name, person.Adress)

	return err
}
