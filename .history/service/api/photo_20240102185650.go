package api

import (
	"Wasa-photo-1905917/service/database"
	"time"
)

type Photo struct {
	Id_photo int64              `json:"id"`
	User_ID  string             `json:"user_id"`
	Likes    []database.Like    `json:"likes"`
	Comments []database.Comment `json:"comments"`
	Date     time.Time          `json:"date"`
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		Id_photo: p.Id_photo,
		User_ID:  p.User_ID,
		Likes:    p.Likes,
		Comments: p.Comments,
		Date:     p.Date,
	}
}
