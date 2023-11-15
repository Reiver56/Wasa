package structures

type Photo struct{
	Id_photo string `json:"id"`
	User_ID User_ID `json:"user"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}