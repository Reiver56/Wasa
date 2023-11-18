package database

import (
	"database/sql"
	"errors"
	
	
)

var query_GETUSER = `SELECT id FROM users WHERE id = ?`

//Database fuction that add a new user in database
func (db *appdbimpl) CreateUser(user User) (User, error) {
	var u User
	u.Nickname = user.Nickname

	//check if exist user id
	row, err := db.c.Query(`SELECT id FROM users`, u.Nickname)
	if err != nil {
		return u, err
	}
	var id string
	for row.Next() {
		if row.Err() != nil{
			return u, err
		}
		err = row.Scan(&id)
		if err != nil && !errors.Is(err, sql.ErrNoRows){
			return user, err
			



	u.ID = user.Nickname
	u.Nickname = user.Nickname


	
}

//Check if user exist in database
func (db *appdbimpl) ExistUser(nickname string) (bool, error) {
	var existUser string
	err := db.c.QueryRow(`SELECT nickname FROM users WHERE nickname = ? `, nickname).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows){
		return false, nil
	}
	return existUser != "",err
	
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
