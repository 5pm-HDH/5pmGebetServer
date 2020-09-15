package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func executeSqlFile(filename string, db *sql.DB) {
	file, _ := ioutil.ReadFile(filename)

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			log.Print(err)
		}
	}
}

func openDatabase(dbPath, key string) *sql.DB {
	// check if the sqlite file exists
	create := true
	if fileExists("./pray.db") {
		create = false
	}

	// open sql connection and database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	if create {
		executeSqlFile("./ddl.sql", db)
		if key != "" {
			log.Println(db.Exec("INSERT INTO authorization(auth_key) VALUES(?)", key))
		} else {
			log.Println("not inserting a master key into an empty database")
		}
	}
	return db
}
