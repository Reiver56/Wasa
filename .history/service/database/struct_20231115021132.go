package database

import (
	"time"
	"Wasa-photo-1905917/service/api/structures"
)


type Like struct{
	ID_like int `json:"id"`
	ID_user int `json:"id_user"`
}
type Comments struct{
	ID_comment int `json:"id"`
	ID_user 
}



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

