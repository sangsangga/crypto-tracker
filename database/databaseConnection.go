package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Client *sql.DB

func StartDatabase() {
	var err error
	Client, err = sql.Open("sqlite3", "coin.db")

	const create string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  email text NOT NULL,
  password TEXT NOT NULL
  ); CREATE TABLE IF NOT EXISTS usercoins (
	id INTEGER NOT NULL PRIMARY KEY,
	coinId TEXT NOT NULL,
	userId INTEGER NOT NULL,
	FOREIGN KEY(userId) REFERENCES users(id)
	);`

	if err != nil {
		log.Fatal(err)
	}

	if _, err := Client.Exec(create); err != nil {
		log.Fatal(err)
	}

	if _, err := Client.Exec("PRAGMA foreign_keys=ON;"); err != nil {
		log.Fatal(err)
	}

}
