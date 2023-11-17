package structures

import "Wasa-photo-1905917/service/database"

type profile struct{
	User      User      `json: "user"`
	Followers database.Profile. `json: "followers"`
	Following Following `json "following"`
	Photos    []Photo   `json: "photos"`

}