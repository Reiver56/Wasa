package structures

import 

type profile struct{
	User      User      		`json: "user"`
	Followers []database.User   `json: "followers"`
	Following []database.User   `json "following"`
	Photos    []database.Photo  `json: "photos"`

}