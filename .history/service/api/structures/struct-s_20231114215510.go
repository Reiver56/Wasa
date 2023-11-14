package structures

//in this file, describe singular  object 

type id_Comment struct{
	ID_comment int `json:"id"`
}

type id_Like struct{
	ID_like int `json:"id"`
}

type Like struct{
	id_Like id_Like `json: "id"`

}

type Comment struct{
	Comment string `json:"comment"`
}

type Followers struct{
	Followers []User `json:"followers"`
}

type Following struct{
	Following []User `json:"following"`
}

