package api

//  in this file, describe singular  object


type Like struct {
	ID_photo int    `json: "id_like"`
	User    string `json: "user"`
}
func (l Like) ToDatabase() database.{
	return database.User_ID{
		ID: u.ID,	
		
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
