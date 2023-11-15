package database

import (
	"time"
)



type User struct{
	ID string `json:"id"`
	Password string `json:"password"`
}

type Photo struct{
	ID string  `json:"id"`
}

