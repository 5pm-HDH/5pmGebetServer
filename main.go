package main

import (
	"5pmGebetServer/web"
	"log"
	"net/http"
	"os"
)

func main() {

	db := openDatabase()
	defer db.Close()
	web.SetDatabase(db)

	//port := flag.String("p", "8080", "port to serve on")
	//directory := flag.String("d", "web/static/", "the directory of static file to host")
	//flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("please specify a port in $PORT")
	}

	masterkey := os.Getenv("MASTERKEY")
	if masterkey == "" {
		log.Fatal("pleas set a master key in $MASTERKEY")
	}
	err := setMasterKey(masterkey, db)
	if err != nil {
		log.Fatalf("error setting master key: %v", err)
	}

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
