package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("sqlite3", "api.db")

    if err != nil {
        panic("Could not connect to database.")
    }

    DB.SetMaxOpenConns(10)
    DB.SetMaxIdleConns(5)

    createTables()
}

func createTables(){

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
		)
		`
		_, err := DB.Exec(createUsersTable)

		if err != nil{
			panic("Could not create user table")
		}


	//using back ticks to create string across multiple lines
		createEventsTable := `
		CREATE TABLE IF NOT EXISTS event (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTIme DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES user(id)
		)
		`
		_, err = DB.Exec(createEventsTable)

		if err != nil{
			panic("Could not create events table")
		}





}
