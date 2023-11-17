package database

import (
	"time"
	"Wasa-photo-1905917/service/api/str"
)



type User struct{
	ID string `json:"id"`
	Password string `json:"password"`
}

type Photo struct{
	Id_photo string      `json:"id"`
	User_ID  string     `json:"user_id"`
	Likes 	 []Like 	 `json: "likes"`
	Comments []Comment   `json: "comments"`
	time 	 time.Time   `json: "date"`
}
