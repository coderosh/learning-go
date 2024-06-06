package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("sqlite3", "short.db")
	DB = db

	if err != nil {
		panic("DB connection failed")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTables()
	if err != nil {
		log.Panicln("Table creation failed", err)
	}
}

func createTables() error {
	createUsersTableQuery := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTableQuery)
	if err != nil {
		return err
	}

	createUrlTableQuery := `CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		long_url TEXT NOT NULL,
		short_code TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createUrlTableQuery)
	if err != nil {
		return err
	}

	createAnalyticsTableQuery := `CREATE TABLE IF NOT EXISTS analytics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		views INTEGER,
		date_time DATETIME NOT NULL,
		user_agent string,
		ip string,
		url_id string
	)`
	_, err = DB.Exec(createAnalyticsTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func PrepareAndExec(query string, params ...any) (int64, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
