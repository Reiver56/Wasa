package api

import "Wasa-photo-1905917/service/database"

//  in this file, describe singular  object


type Like struct {
	ID_photo int    `json: "id_like"`
	User    string `json: "user"`
}
func (l Like) ToDatabase() database.Like{
	return database.Like{
		ID: l.ID_photo,
		User: l.User,
	}	
	}


type Comment struct {
	ID_Comment int    `json: "id_comment"`
	Comment    string `json:"comment"`
}

type Followers struct {
	Followers []User `json:"followers"`
}

type Following struct {
	Following []User `json:"following"`
}
