package structures

type Photo struct{
	Id_photo string `json:"id"`
	Id_user string  `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment 
}