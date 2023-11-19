package database

import (
	"database/sql"
	"errors"
	"fmt"
)

//Database fuction that add a new user in database
func (db *appdbimpl) CreateUser(u User) error {
	