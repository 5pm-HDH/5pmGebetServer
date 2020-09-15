package main

import (
	"5pmGebetServer/web"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("using default port: %s", port)
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./pray.db"
		log.Printf("using default db path: %s", dbPath)
	}

	masterkey := os.Getenv("MASTERKEY")
	if masterkey == "" {
		log.Println("no new master key specified")
	}

	db := openDatabase(dbPath, masterkey)
	defer db.Close()
	web.SetDatabase(db)

	directory := os.Getenv("DIRECTORY")
	if directory == "" {
		log.Println("using default directory \"web/static/\"")
		directory = "web/static/"
	}

	http.HandleFunc("/api", web.ApiHandle)
	http.HandleFunc("/api/view", web.ApiView)
	http.HandleFunc("/api/key", web.ApiKey)
	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
