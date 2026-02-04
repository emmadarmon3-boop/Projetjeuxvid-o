package jeux

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	data_sql := "./basedonnee.sqlite"

	sqlite, err := sql.Open("sqlite3", data_sql)

	if err != nil {
		log.Fatal(err)
	}

	defer sqlite.Close()

	create_user_table :=
		`CREATE TABLE if not exists users 
			(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE
			);

		`

	_, err = sqlite.Exec(create_user_table)
	if err != nil {
		log.Fatal(err)
	}

	create_games_table :=
		`CREATE TABLE if not exists games 
			(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			platform TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
			);
		`

	_, err = sqlite.Exec(create_games_table)
	if err != nil {
		log.Fatal(err)
	}
}

func AddTable(pseudo string) {
	data_sql := "./basedonnee.sqlite"

	sqlite, err := sql.Open("sqlite3", data_sql)

	if err != nil {
		log.Fatal(err)
	}

	defer sqlite.Close()

	_, err = sqlite.Exec(`INSERT INTO users(username) 
		VALUES (?);`, pseudo)

	if err != nil {
		log.Fatal(err)
	}
}
