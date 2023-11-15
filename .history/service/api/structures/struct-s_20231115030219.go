package structures

//in this file, describe singular  object 

type id_Comment struct{
	ID_comment int `json:"id"`
}

type id_Like struct{
	ID_like int `json:"id"`
}

type Like struct{
	ID_Like int `json: "id_like"`
	User 	string 	`json: "user"`

}

type Comment struct{
	ID_Comment id_Comment `json: "id_comment"`
	Comment string `json:"comment"`
}

type Followers struct{
	Followers []User `json:"followers"`
}

type Following struct{
	Following []User `json:"following"`
}

