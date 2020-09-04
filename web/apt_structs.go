package web

import "time"

type Prayer struct {
	Id       string    `json:"id"`
	Content  string    `json:"prayer"`
	Public   bool      `json:"is_public"`
	Approved bool      `json:"approved"`
	Date     time.Time `json:"date"`
}

var JsonToSql = map[string]string{
	"id":        "id",
	"prayer":    "prayer_text",
	"is_public": "public",
	"created":   "date",
	"approved":  "approved",
}
