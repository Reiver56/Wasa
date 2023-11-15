package structures

import "Wasa-photo-1905917/service/database"

type profile struct{
	User      User      `json: "user"`
	Followers datab   `json: "followers"`
	Following []User    `json "following"`
	Photos    []Photo   `json: "photos"`

}