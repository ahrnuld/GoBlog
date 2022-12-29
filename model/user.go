package model

import "time"

type User struct {
	Id        int    `json:"firstname"`
	Username  string `json:"lastname"`
	Password  string
	CreatedAt time.Time
}
