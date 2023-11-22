package database

import (
	"database/sql"
	"errors"
	"fmt"
)

//  Database fuction that add a new user in database
func (db *appdbimpl) CreateUser(u User) error {

	//  -------INSERT USER IN DATABASE-------
	_, err := db.c.Exec(`INSERT INTO users (id, nickname) VALUES(?, ?)`, u.ID, u.Nickname)
	if err != nil {
		fmt.Println("Error inserting user in database")
		return err
	}
	return nil

}

//  Check if user exist in database
func (db *appdbimpl) ExistUser(nickname string) (bool, error) {
	var existUser string
	err := db.c.QueryRow(`SELECT nickname FROM users WHERE nickname = ?`, nickname).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existUser != "", err

}

var query_GETUSER = `SELECT nickname FROM users WHERE nickname = ?`

//  Get user object from database
func (db *appdbimpl) GetUser(username string) (User, error) {
	var user User
	err := db.c.QueryRow(query_GETUSER, username).Scan(&user.Nickname)
	return user, err
}

//  Fuction that modifies a user's username
func (db *appdbimpl) SetNewUsername(userID string, username string) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id = ?`, username, userID)
	fmt.Println(err)
	return err

}

func (db *appdbimpl) CheckUser(user User) (bool) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id_user = ?",
		user.ID)

	if err != nil {
		//   Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true
	}

	//   If the counter is 1 then the user exists
	if cnt == 1 {
		return true
	}
	return false, nil
}
