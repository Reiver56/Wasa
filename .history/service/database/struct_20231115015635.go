package database

import (
	"time"
)

type User struct{
	ID int `json:"id"`
	PasswordHash string `json:"password_hash"`
}