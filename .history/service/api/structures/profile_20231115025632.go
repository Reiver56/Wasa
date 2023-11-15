package structures

import "Wasa-photo-1905917/service/database"

type Profile struct{
	User_id     string      		`json: "user"`
	Followers []database.User   `json: "followers"`
	Following []database.User   `json "following"`
	Photos    []database.Photo  `json: "photos"`

}

func (p *Profile) ToDatabase() database.Profile{
	return database.Profile{
		ID: p.User_id,
		Followers: ,
		
	}
}