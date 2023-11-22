package database

//  insert ban in the database, with the banner and the banned
func (db *appdbimpl) BanUser(banner User, banned User) error {
	if (ch
	_, err := db.c.Exec("INSERT INTO banned (banner,banned) VALUES (?,?)", banned.ID, banner.ID)
	if err != nil {
		return err
	}
	return nil
}
//  delete ban in the database, with the banner and the banned
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE (banner = ? AND banned = ?)", banned.ID, banner.ID)
	if err != nil {
		return err
	}
	return nil
}