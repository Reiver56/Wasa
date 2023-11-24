package database
import (
	"database/sql"
	"errors"
)

var query_FINDUSER = `SELECT ID AND password FROM User WHERE ID = ? AND password = ?`

//  Database fuction that add a new user in database 
func (db *appdbimpl) CreateUser(username string, password string) error{
	_, err := db.c.Exec("INSERT INTO users (username,password) VALUES (?,?), username, password")

	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) ExistUser(username string, password string) (bool, error){
	var existUser string
	err := db.c.QueryRow(query_FINDUSER, username, password).Scan(&existUser)
	if err != nil && errors.Is(err,sql.ErrNoRows){
		return false, err
		}
	return existUser != "",err	 
}