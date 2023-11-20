package database

func (db *appdbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "INSERT banned (banner, banned) VALUES (?,?)", banned.ID)
	if err != nil {
		return err
	}
	return nil
}