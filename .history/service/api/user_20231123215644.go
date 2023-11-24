package api

import (
	"Wasa-photo-1905917/service/database"
	"regexp"
)

//define structure of user object with all information

type User struct {
	ID        string `json:"id"`
	Nickname string `json:"nickname"`
}

//covert user in api package in database package
func (u *User) ToDatabase() database.User{
	return database.User{
		ID: u.ID,
		Nickname: u.Nickname,
		
	}
}

func (u _User) ToDatabase() database.User_ID{
	return database.User_ID{
		ID: u.ID,	
		
	}
}


func (u *User) FromDatabase(User database.User) error{
	u.ID = User.ID
	u.Nickname = User.Nickname
	
	return nil

}


// check identifier and password with regex
func (u *User) isValidID() bool{
	identifier := u.Nickname
	
	
	valid := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	if(valid.MatchString(identifier)){
		return true
		}else{return false} 
}






