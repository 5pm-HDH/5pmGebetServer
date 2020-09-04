package web

import (
	"database/sql"
	"net/http"
)

var database *sql.DB

func SetDatabase(conn *sql.DB) {
	database = conn
}

func apiGet(w http.ResponseWriter, r *http.Request) {

}

func apiPost(w http.ResponseWriter, r *http.Request) {

}

func apiPut(w http.ResponseWriter, r *http.Request) {

}

func ApiHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		apiGet(w, r)
		return
	case "POST":
		apiPost(w, r)
		return
	case "PUT":
		apiPut(w, r)
		return
	}
}
