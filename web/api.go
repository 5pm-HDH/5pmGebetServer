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

// try to find this auth_key in the corresponding type table
// true if the key is valid
// false if no match was found
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

//API function to insert a new prayer into the database
func apiPost(w http.ResponseWriter, r *http.Request) {

	//decode json request
	dec := json.NewDecoder(r.Body)
	var p Prayer
	err := dec.Decode(&p)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}
	p.Date = time.Now()

	//insert into database
	err = DatabaseInsertPrayer(p)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

//API function to insert a new prayer into the database
func apiPut(w http.ResponseWriter, r *http.Request) {

	//authenticate with key in QUERY
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(1, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	//decode json request
	dec := json.NewDecoder(r.Body)
	var prayer map[string]interface{}
	err := dec.Decode(&prayer)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	//begin change transaction
	tx, err := database.Begin()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	// id attribute must be given
	id, _ := strconv.Atoi(prayer["id"].(string))
	// update the value in the database for each other attribute in the request
	for key, value := range prayer {
		if key == "id" {
			continue
		}

		// map only valid json keys to sql keys
		// this will fail if the users puts any other string here
		// -> NO SQL INJECTION POSSIBLE
		sqlKey := JsonToSql[key]

		query := "UPDATE prayer SET % = ? WHERE id = ?"
		query = strings.Replace(query, "%", sqlKey, 1)

		//update one attribute
		_, err = tx.Exec(query, value, id)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			_ = tx.Rollback()
			return
		}

	}

	//commit database transaction
	err = tx.Commit()
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

}

//handle switch for /api to distinguish between Methods (GET,POST,PUT)
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

//View API to only return reduced filtered and approved data
func ApiView(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(2, key) {
		w.WriteHeader(401)
		_, _ = io.WriteString(w, "401 Unauthorized")
		return
	}

	// optional argument to filter
	// only get prayer with id's bigger than <l>
	l := r.URL.Query().Get("l")
	i, err := strconv.Atoi(l)
	if err != nil {
		i = 0
	}

	//database request
	prayers, err := DatabaseGetPrayerTodayPublicApproved(i)
	//only PlayerSlim(reduced) prayers are returned
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	//encode response
	enc := json.NewEncoder(w)
	err = enc.Encode(prayers)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
}

//Key management API
func ApiKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if !apiAuthenticate(0, key) {
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
