package structures

type Photo struct{
	Id_photo string `json:"id"`
	uUser   `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}