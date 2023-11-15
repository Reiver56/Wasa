package structures

import "Wasa-photo-1905917/service/database"

type profile struct{
	User      User      		`json: "user"`
	Followers []database.User   `json: "followers"`
	Following []database.User   `json "following"`
	Photos    []dat   		`json: "photos"`

}