package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

var schema = `
CREATE TABLE posts (
    id integer primary key autoincrement,
    title varchar(255),
    author varchar(255),
    content text
);
`

func ConnectDatabase() {
	var err error
	// exactly the same as the built-in
	database, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		log.Fatalln(err)
	}
	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	database.MustExec(schema)
	db = database
}

func GetDB() *sqlx.DB {
	return db
}
