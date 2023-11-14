package structures

import(
	"regexp"
)

//define structure of user object with all information

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}

func (u *User) isValidID() bool{
	identifier := u.ID
	password := u.Password
	valid := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	if(valid.MatchString(identifier) && valid.MatchString(password) == true){return true}else{} 


}
