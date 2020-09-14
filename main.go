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
		log.Fatal("please specify a port in $PORT")
	}

	masterkey := os.Getenv("MASTERKEY")
	if masterkey == "" {
		log.Fatal("pleas set a master key in $MASTERKEY")
	}

	db := openDatabase(masterkey)
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
