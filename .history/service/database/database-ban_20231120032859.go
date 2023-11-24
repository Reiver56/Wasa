package database

//  insert ban in the database, with the banner and the banned
func (db *appdbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.Exec("INSERT banned (banner, banned) VALUES (?,?)", banned.ID)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec(db.ctx, "INSERT banned (banner, banned) VALUES (?,?)", banned.ID)
	if err != nil {
		return err
	}
	return nil
}