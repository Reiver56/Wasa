package database

import(

)

var query_FINDUSER = `SELECT ID AND password FROM User WHERE ID = ? AND password = ?`

func (db *appdbimpl) existUser(username string, password string) (bool, error){
	var exists
}