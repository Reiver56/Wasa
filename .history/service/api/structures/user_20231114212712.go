package structures

import(
	"regexp"
)

//define structure of user object with all information

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}

func (u *User) validID() bool{
	identifier := u.ID
	password := u.Password


}
