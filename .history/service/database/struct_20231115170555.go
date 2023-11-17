package database

import (
	"time"
)

type Like struct {
	ID_like int    `json:"id"`
	ID_user string `json:"id_user"`
}

type Comment struct {
	ID_comment   int    `json:"id"`
	ID_user      string `json:"is_user"`
	Text_comment string `json:"text"`
}

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
type Profile struct {
	ID        string  `json:"id"`
	Followers []User  `json: "followers"`
	Following []User  `json: "following"`
	Photos    []Photo `json: "photos"`
}
type Photo struct {
	Id_photo string    `json:"id"`
	User_ID  string    `json:"user_id"`
	Likes    []Like    `json: "likes"`
	Comments []Comment `json: "comments"`
	Date     time.Time `json: "date"`
}
