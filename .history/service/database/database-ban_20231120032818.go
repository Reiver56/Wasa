package database

//  insert ban in the database, with the banner and the banned
func (db *appdbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "INSERT banned (banner, banned) VALUES (?,?)", banned.ID)
	if err != nil {
		return err
	}
	return nil
}
func (db *appdbimpl) UnanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "INSERT banned (banner, banned) VALUES (?,?)", banned.ID)
	if err != nil {
		return err
	}
	return nil
}