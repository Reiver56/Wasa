package structures

import(
	"regexp"
)

//define structure of user object with all information

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}

func (u *User) isvalidID() bool{
	identifier := u.ID
	password := u.Password
	validUser := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	return validUser.MatchString(identifier)


}
