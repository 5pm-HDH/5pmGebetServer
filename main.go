package main

import (
	"5pmGebetServer/web"
	"flag"
	"log"
	"net/http"
)

func main() {

	db := openDatabase()
	defer db.Close()
	web.SetDatabase(db)

	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", "web/static/", "the directory of static file to host")
	flag.Parse()

	http.HandleFunc("/api", web.ApiHandle)
	http.HandleFunc("/api/beamer", web.ApiBeamer)
	http.HandleFunc("/api/key", web.ApiKey)
	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
