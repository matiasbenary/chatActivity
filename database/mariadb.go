package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDBMaria() *sql.DB {
	db, err := sql.Open("mysql", "homestead:secret@tcp(localhost:3306)/fonselp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := `	
	CREATE TABLE IF NOT EXISTS room (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		private TINYINT NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("%q: %s\n", err, sqlStmt)
	}

	sqlStmt = `	
	CREATE TABLE IF NOT EXISTS user (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255)  NULL,
		role_id VARCHAR(255)  NULL
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("%q: %s\n", err, sqlStmt)
	}

	sqlStmt = `	
	CREATE TABLE IF NOT EXISTS message (
		id VARCHAR(255) NOT NULL PRIMARY KEY,
		value VARCHAR(255) NOT NULL,
		send_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_id VARCHAR(255) NOT NULL ,
		room_id VARCHAR(255) NOT NULL 
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal("%q: %s\n", err, sqlStmt)
	}

	return db
}
