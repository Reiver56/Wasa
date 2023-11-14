package structures

import()

//define structure of user with all information
type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
}
