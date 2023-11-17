package database

import(

)

var query_FINDUSER = `SELECT ID AND password FROM User WHERE ID = ? AND `

func (db *appdbimpl) existUser(username string)