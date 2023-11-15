package structures

type Photo struct{
	Id_photo string `json:"id"`
	User   u.ID  `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}