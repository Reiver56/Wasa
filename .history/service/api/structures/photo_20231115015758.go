package structures

type Photo struct{
	Id_photo string `json:"id"`
	User   IdUser  `json:"userId"`
	Likes []Like 	`json: "likes"`
	Comments []Comment `json: "comments"`
}