package web

import (
	"log"
	"time"
)

//get all prayers of today from the database
func DatabaseGetPrayerToday() ([]Prayer, error) {
	//query the database
	yesterday := time.Now().Add(-24 * time.Hour).Unix()
	rows, err := database.Query("SELECT id,prayer_text,approved,public,created FROM prayer WHERE created > ? ", yesterday)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//convert rows to Prayer struct
	var prayers []Prayer
	for rows.Next() {
		var p Prayer
		err := rows.Scan(&p.Id, &p.Content, &p.Approved, &p.Public, &p.Date)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		prayers = append(prayers, p)
	}

	return prayers, nil
}

//insert prayer into the database
func DatabaseInsertPrayer(p Prayer) error {
	// begin transaction
	tx, _ := database.Begin()
	// insert
	_, err := tx.Exec("INSERT INTO prayer(prayer_text, public, created, approved, state, original_test) VALUES (?,?,?,0,0,?)",
		p.Content, p.Public, time.Now().Unix(), p.Content)

	// on error rollback
	if err != nil {
		_ = tx.Rollback()
		log.Println(err)
		return err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Get all PrayerSlim of today in the database (BEAMER VIEW)
func DatabaseGetPrayerTodayPublicApproved(lastId int) ([]PrayerSlim, error) {
	//query the database
	yesterday := time.Now().Add(-24 * time.Hour).Unix()
	rows, err := database.Query("SELECT id,prayer_text FROM prayer WHERE approved = 1 AND public = 1 AND created > ? AND id > ? ORDER BY id", yesterday, lastId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//convert rows to Prayer struct
	var prayers []PrayerSlim
	for rows.Next() {
		var p PrayerSlim
		err := rows.Scan(&p.Id, &p.Content)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		prayers = append(prayers, p)
	}

	return prayers, nil
}

//get all auth_keys of a given type/category
func DatabaseGetKeys(keyType int) []string {
	var keys []string
	rows, err := database.Query("SELECT auth_key FROM authorization WHERE type = 0 OR type = ? ", keyType)
	if err != nil {
		log.Println(err)
		return keys
	}

	for rows.Next() {
		var p string
		err := rows.Scan(&p)
		if err != nil {
			log.Println(err)
			return keys
		}
		keys = append(keys, p)
	}

	return keys
}

//get all auth_keys in the database.
func DatabaseGetAllKeys() []AuthKey {
	var keys []AuthKey
	rows, err := database.Query("SELECT * FROM authorization")
	if err != nil {
		log.Println(err)
		return keys
	}

	for rows.Next() {
		var k AuthKey
		err := rows.Scan(&k.Key, &k.Type)
		if err != nil {
			log.Println(err)
			return keys
		}
		keys = append(keys, k)
	}

	return keys
}
