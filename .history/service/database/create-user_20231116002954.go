package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CreateUser(username string, password string) error{
	_, err := db.c.Exec()
}