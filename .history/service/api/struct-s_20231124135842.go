package api

import "Wasa-photo-1905917/service/database"

//in this file, describe singular  object


type Like struct {
	ID_photo int64    `json:"id_like"`
	User    string `json:"id"`
}
func (l Like) ToDatabase() database.Like{
	return database.Like{
		ID_photo: l.ID_photo,
		ID_user: l.User,
	}	
	}
type Comment struct {
	ID_comment   int64    `json:"id"`
	ID_user      string `json:"id_user"`
	ID_photo     int64    `json:"id_photo"`
	Text_comment string `json:"text"`
}
func (c Comment) ToDatabase() database.Comment{
	return database.Comment{
		ID_comment: c.ID_comment,
		ID_user: c.ID_user,
		ID_photo: c.ID_photo,
		Text_comment: c.Text_comment,
	}	
	}


type Followers struct {
	Followers []User `json:"followers"`
}

type Following struct {
	Following []User `json:"following"`
}
