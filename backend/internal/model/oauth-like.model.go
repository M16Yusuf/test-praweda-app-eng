package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginBody struct {
	Username string `json:"username"`
}
