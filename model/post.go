package model

import "time"

type Post struct {
	Id       int
	Title    string
	PostedAt time.Time
	Content  string
}
