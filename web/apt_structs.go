package web

import "time"

type Prayer struct {
	Id       string    `json:"id"`
	Content  string    `json:"prayer"`
	Public   bool      `json:"is_public"`
	Approved bool      `json:"approved"`
	Date     time.Time `json:"date"`
}

//reduced Prayer that only contains the text and id
type PrayerSlim struct {
	Id      string `json:"id"`
	Content string `json:"prayer"`
}

type AuthKey struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

//lookup table to convert json tags to SQL columns
var JsonToSql = map[string]string{
	"id":        "id",
	"prayer":    "prayer_text",
	"is_public": "public",
	"created":   "date",
	"approved":  "approved",
}
