package structures

type Photo struct{
	Id_photo string `json:"id"`
	User   id.User  `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}