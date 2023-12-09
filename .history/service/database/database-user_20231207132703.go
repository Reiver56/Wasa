package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Database fuction that add a new user in database
func (db *appdbimpl) CreateUser(u User) error {
	var user User
	user.Nickname = u.Nickname
	user.ID = user.Nickname
	fmt.
	// -------INSERT USER IN DATABASE-------
	_, err := db.c.Exec(`INSERT INTO users (id, nickname) VALUES(?, ?)`, u.ID, u.Nickname)
	if err != nil {
		return err
	}
	return nil

}

// Check if user exist in database
func (db *appdbimpl) ExistUser(nickname string) (bool, error) {
	var existUser string
	err := db.c.QueryRow(`SELECT nickname FROM users WHERE nickname = ?`, nickname).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existUser != "", err

}

var query_GETUSER = `SELECT nickname FROM users WHERE nickname = ?`

// Get user object from database
func (db *appdbimpl) GetUser(username string) (User, error) {
	var user User
	err := db.c.QueryRow(query_GETUSER, username).Scan(&user.Nickname)
	return user, err
}
func (db *appdbimpl) GetNickname(id string) (string, error) {
	var nickname string
	err := db.c.QueryRow(`SELECT nickname FROM users WHERE id = ?`, id).Scan(&nickname)
	return nickname, err
}

// Fuction that modifies a user's username
func (db *appdbimpl) SetNewUsername(userID string, username string) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id = ?`, username, userID)
	return err

}

func (db *appdbimpl) CheckUser(user User) bool {
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id_user = ?", user.ID)
	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return false
	}
	return true
}

func (db *appdbimpl) GetMyStream(user User) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photo WHERE user_id IN (SELECT followed FROM follow WHERE follower = ?) ORDER BY date DESC", user.ID)
	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.Id_photo, &photo.User_ID, &photo.Likes, &photo.Comments, &photo.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return photos, nil
}
