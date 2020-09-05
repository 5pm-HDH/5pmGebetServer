package main

import (
	"5pmGebetServer/web"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
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
			fmt.Print(err)
		}
	}
}

func openDatabase() *sql.DB {
	// check if the sqlite file exists
	create := true
	if fileExists("./pray.db") {
		create = false
	}

	// open sql connection and database
	db, err := sql.Open("sqlite3", "./pray.db")
	if err != nil {
		log.Fatal(err)
	}

	if create {
		executeSqlFile("./ddl.sql", db)
	}
	return db
}

func main() {

	db := openDatabase()
	defer db.Close()
	web.SetDatabase(db)

	port := flag.String("p", "80", "port to serve on")
	directory := flag.String("d", "web/static/", "the directory of static file to host")
	flag.Parse()

	http.HandleFunc("/api", web.ApiHandle)
	http.HandleFunc("/api/beamer", web.ApiBeamer)
	http.HandleFunc("/api/key", web.ApiKey)
	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
