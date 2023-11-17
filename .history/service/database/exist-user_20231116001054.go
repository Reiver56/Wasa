package database

import (
	"database/sql"
	"errors"
)

var query_FINDUSER = `SELECT ID AND password FROM User WHERE ID = ? AND password = ?`

func (db *appdbimpl) existUser(username string, password string) (bool, error){
	var existUser string
	err := db.c.QueryRow(query_FINDUSER, username, password).Scan(&existUser)
	if err != nil {
		return false, err
		}
	return existUser != "",err	 
}