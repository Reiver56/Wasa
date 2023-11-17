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
	ID_user      string `json:"id_user"`
	Text_comment string `json:"text"`
}
type User_ID struct{
	ID string `json:"id"`
}
type nickname struct{
	nickname string `json:"nickname"`
}

type User struct {
	ID string `json:"id"`
	nickname string `json:"nickname"`
	
}
type Profile struct {
	ID        string  `json:"id"`
	Nickname  string  `json:"nickname"`
	Followers []User  `json: "followers"`
	Following []User  `json: "following"`
	Photos    []Photo `json: "photos"`
}
type Photo struct {
	Id_photo string    `json:"id"`
	ni
	User_ID  string    `json:"user_id"`
	Likes    []Like    `json: "likes"`
	Comments []Comment `json: "comments"`
	Date     time.Time `json: "date"`
}
