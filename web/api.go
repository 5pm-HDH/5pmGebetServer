package web

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var database *sql.DB

func SetDatabase(conn *sql.DB) {
	database = conn
}

func apiAuthenticate(accType int, key string) bool {
	keys := DatabaseGetKeys(accType)

	for _, k := range keys {
		if k == key {
			return true
		}
	}

	return false
}

/*
	API function to get all prayers of today!
	authentication with query key
*/
func apiGet(w http.ResponseWriter, r *http.Request) {
	//authenticate with key in QUERY
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(1, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	//get to days prayer from database
	prayers, err := DatabaseGetPrayerToday()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(prayers)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

func apiPost(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	var p Prayer
	err := dec.Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}
	p.Date = time.Now()

	err = DatabaseInsertPrayer(p)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

func apiPut(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query().Get("key")
	if !apiAuthenticate(1, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	dec := json.NewDecoder(r.Body)
	var prayer map[string]interface{}
	err := dec.Decode(&prayer)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	tx, err := database.Begin()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	id := int(prayer["id"].(float64))
	for key, value := range prayer {
		if key == "id" {
			continue
		}

		sqlKey := JsonToSql[key]
		query := "UPDATE prayer SET % = ? WHERE id = ?"
		query = strings.Replace(query, "%", sqlKey, 1)

		_, err = tx.Exec(query, value, id)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			_ = tx.Rollback()
			return
		}

	}

	err = tx.Commit()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

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

func ApiBeamer(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(2, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	l := r.URL.Query().Get("l")
	i, err := strconv.Atoi(l)
	if err != nil {
		i = 0
	}

	prayers, err := DatabaseGetPrayerTodayPublicApproved(i)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(prayers)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

func ApiKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(2, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	switch r.Method {
	case "PUT":
		// DELETE an AUTH KEY
		dec := json.NewDecoder(r.Body)
		var k AuthKey
		err := dec.Decode(&k)
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}
		_, err = database.Exec("DELETE FROM authorization WHERE auth_key = ?", k.Key)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

	case "GET":

		keys := DatabaseGetAllKeys()
		enc := json.NewEncoder(w)
		err := enc.Encode(keys)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}

	case "POST":
		// ADD an AUTH KEY
		dec := json.NewDecoder(r.Body)
		var k AuthKey
		err := dec.Decode(&k)
		if err != nil {
			w.WriteHeader(400)
			log.Println(err)
			return
		}

		_, err = database.Exec("INSERT INTO authorization(auth_key, type) VALUES (?,?)", k.Key, k.Type)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
	}
}
