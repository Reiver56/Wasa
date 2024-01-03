package database

import (
	"errors"
)
/*
GetBans returns the list of users banned by the user with the given ID.
*/
func (db *appdbimpl) GetBans(bannerID string) ([]User, error){
	var bans []User
	rows, err := db.c.Query("SELECT banned FROM ban WHERE banner = ? ", bannerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() != nil {
			

}

// insert ban in the database, with the banner and the banned
func (db *appdbimpl) BanUser(banner User, banned User) error {
	if db.CheckUser(banner) {
		_, err := db.c.Exec("INSERT INTO banned (banner,banned) VALUES (?,?)", banned.ID, banner.ID)
		if err != nil {
			return err
		}
		return nil
	}
	err := errors.New("User not found")
	return err
}

// delete ban in the database, with the banner and the banned
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE banner = ? AND banned = ?", banned.ID, banner.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) IsBanned(id_banned string, id_banner string) bool {

	var banned string
	err := db.c.QueryRow("SELECT banned FROM banned WHERE banner = ? AND banned = ?", id_banner, id_banned).Scan(&banned)
	if err == nil {
		return true
	}
	return false

}
