package web

import "time"

type Prayer struct {
	Id       string    `json:"id"`
	Content  string    `json:"prayer"`
	Public   bool      `json:"is_public"`
	Approved bool      `json:"approved"`
	Date     time.Time `json:"date"`
}
