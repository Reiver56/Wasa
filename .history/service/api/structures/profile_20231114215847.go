package structures

import(
	
)

type profile struct{
	User      User      `json: "user"`
	Followers Followers `json: "followers"`
	Following Following `json "following"`
	Photos

}