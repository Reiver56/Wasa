package structures

import (
	"Wasa-photo-1905917/service/database"
	"regexp"
)

//define structure of user object with all information
type User_ID struct{
	ID string `json:"id"`
}
type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}



// check identifier and password with regex
func (u *User) isValidID() bool{
	identifier := u.ID
	password := u.Password
	valid := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	if(valid.MatchString(identifier) && valid.MatchString(password) == true){
		return true
		}else{return false} 


}
//covert user in api package in database package
func (u *User) ToDatabase() database.User{
	return database.User{
		User_ID: u.,

	}
}



