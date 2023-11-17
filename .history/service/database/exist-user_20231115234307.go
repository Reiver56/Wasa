package database

import(

)

var query_FINDUSER = `SELECT username AND pass FROM User WHERE ID = ?`

func (db *appdbimpl) existUser(username string)