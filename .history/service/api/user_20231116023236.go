package api

import (
	"Wasa-photo-1905917/service/database"
	"regexp"
)

//  define structure of user object with all information
type User_ID struct{
	ID string `json:"id"`
}

type _User struct{
	ID string `json:"id"`
}
type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}

//  covert user in api package in database package
func (u *User) ToDatabase() database.User{
	return database.User{
		ID: u.ID,
		Password: u.Password,
	}
}
func (u User_ID) ToDatabase() database.ID{
	return database.User_ID{
		ID: u.ID,	
		
	}
}

func (u _User) ToDatabase() database.ID{
	return database.ID{
		ID: u.ID,	
		
	}
}


func (u *User) FromDatabase(User database.User) error{
	u.ID = User.ID
	u.Password = User.Password
	return nil

}


//   check identifier and password with regex
func (u *User) isValidID() bool{
	identifier := u.ID
	password := u.Password
	valid := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	if(valid.MatchString(identifier) && valid.MatchString(password) == true){
		return true
		}else{return false} 
}




