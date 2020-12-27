package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"reflect"
)

func RunSQLite(file string) {
	db, err := sql.Open("sqlite3", "simple.sqlite")
	if err != nil {
		log.Panic(err)
	}

	sqlStmt := "CREATE TABLE IF NOT EXISTS data(id TEXT not null primary key, content TEXT)"
	result, err := db.Exec(sqlStmt)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(reflect.TypeOf(result))
}
