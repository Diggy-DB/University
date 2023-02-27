package auth

import "time"

type Account struct {
	Id         string    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}
