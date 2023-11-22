package database
import("errors")
// insert ban in the database, with the banner and the banned
func (db *appdbimpl) BanUser(banner User, banned User) error {
	if (db.CheckUser(banner) == true) {
	_, err := db.c.Exec("INSERT INTO banned (banner,banned) VALUES (?,?)", banned.ID, banner.ID)
	if err != nil {
		return err
	}
	return nil
	}
	err := errors.New("User not found")
	return err
}
//delete ban in the database, with the banner and the banned
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE banner = ? AND banned = ?", banned.ID, banner.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) IsBanned(id_banned string, id_banner string ) (error,bool) {

	var banned string
	err := db.c.QueryRow("SELECT banned FROM banned WHERE banner = ? AND banned = ?", id_banner, id_banned).Scan(&banned)
	if err != nil {
		return err,false
	}
	return nil,true
	
}