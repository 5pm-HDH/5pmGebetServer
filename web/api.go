package web

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

var database *sql.DB

func SetDatabase(conn *sql.DB) {
	database = conn
}

func apiGet(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("SELECT id,prayer_text,approved,public,created FROM prayer WHERE created = ?", time.Now())
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var prayers []Prayer
	for rows.Next() {
		var p Prayer
		err := rows.Scan(&p.Id, &p.Content, &p.Approved, &p.Public, &p.Date)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		prayers = append(prayers, p)
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(prayers)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)

}

func apiPost(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	var p Prayer
	err := dec.Decode(&p)

	p.Date = time.Now()

	if err != nil {
		w.WriteHeader(400)
		return
	}

	tx, _ := database.Begin()
	_, err = tx.Exec("INSERT INTO prayer(prayer_text, public) VALUES (?,?)", p.Content, p.Public)

	if err != nil {
		w.WriteHeader(400)
		_ = tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
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
