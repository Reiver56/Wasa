package structures

import "Wasa-photo-1905917/service/database"

type profile struct{
	User      User      `json: "user"`
	Followers []User `json: "followers"`
	Following  `json "following"`
	Photos    []Photo   `json: "photos"`

}