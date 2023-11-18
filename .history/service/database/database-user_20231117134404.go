package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var query_GETUSER = `SELECT id FROM User WHERE id = ?`

//Database fuction that add a new user in database
func (db *appdbimpl) CreateUser(username string) error {

	_, err := db.c.Exec("INSERT INTO Users (id) VALUES (?,?), id")
	if err != nil {

		return err
	}
	return nil
}

//Check if user exist in database
func (db *appdbimpl) ExistUser(username string) (bool, error) {
	var existUser string
	err := db.c.QueryRow(`SELECT username FROM users WHERE id = ? `, username).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows){
		return false, nil
	}
	
}

//Get user object from database

func (db *appdbimpl) GetUser(username string) (User, error) {
	var user User
	err := db.c.QueryRow(query_GETUSER, username).Scan(&user.ID)
	return user, err
}

//Fuction that modifies a user's username
func (db *appdbimpl) SetNewUsername(user User_ID, username string) error {

	_, err := db.c.Exec(`UPDATE User SET ID=?, user.ID`)
	if err != nil {

		return err
	}
	return nil
}
