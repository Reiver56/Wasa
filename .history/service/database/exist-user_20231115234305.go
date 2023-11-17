package database

import(

)

var query_FINDUSER = `SELECT username AND passw FROM User WHERE ID = ?`

func (db *appdbimpl) existUser(username string)