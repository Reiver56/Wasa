package database

import(

)

var query_FINDUSER = `SELECT username FROM User WHERE ID = ?`

func (db *appdbimpl) existUser(username)