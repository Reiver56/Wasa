package structures

//in this file, describe singular  object 

type id_Comment struct{
	ID int `json:"id"`
}
type id_Like struct{
	
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

