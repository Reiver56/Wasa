package structures

type Photo struct{
	Id_photo string `json:"id"`
	u string  `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}