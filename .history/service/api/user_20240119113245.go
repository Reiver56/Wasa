package api

import (
	"Wasa-photo-1905917/service/database"
	"regexp"
)

// define structure of user object with all information

type User struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

// covert user in api package in database package
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:       u.ID,
		Nickname: u.Nickname,
	}
}

func (u *User) FromDatabase(User database.User) error {
	// convert user in database package in api package
	u.ID = User.ID
	u.Nickname = User.Nickname
	return nil

}

// check identifier and password with regex
func (u *User) isValidID() bool {
	nickname := u.Nickname
	valid := regexp.MustCompile(`^[a-z][a-z0-9]{2,13}$`)
	if valid.MatchString(nickname) {
		return true
	}
	return false
}
