package database

import (
	"database/sql"
	"errors"
)

var query_FINDUSER = `SELECT ID AND password FROM User WHERE ID = ? AND password = ?`
var query_GETUSER = `SELECT ID AND password FROM User WHERE username = ?`


//  Database fuction that add a new user in database 
func (db *appdbimpl) CreateUser(username string, password string) (User, error){
	_, err := db.c.Exec("INSERT INTO users (username,password) VALUES (?,?), username, password")
	if err != nil {
		return nil,err
	}
	var user User 
	user.ID = username
	user.Password = password
	return user, nil
}
//  Check if user exist in database
func (db *appdbimpl) ExistUser(username string, password string) (bool, error){
	var existUser string
	err := db.c.QueryRow(query_FINDUSER, username, password).Scan(&existUser)
	if err != nil && errors.Is(err,sql.ErrNoRows){
		return false, err
		}
	return existUser != "",err	 
}
//  Get user object from database

func (db *appdbimpl) GetUser(username string, password string) (User, error){
	var user User
	err := db.c.QueryRow(query_GETUSER, username).Scan(&user.ID, &user.Password)
	return user, err
}