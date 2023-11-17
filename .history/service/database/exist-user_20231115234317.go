package database

import(

)

var query_FINDUSER = `SELECT ID AND pass FROM User WHERE ID = ?`

func (db *appdbimpl) existUser(username string)