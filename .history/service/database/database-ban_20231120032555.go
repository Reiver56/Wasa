package database

func (db *appbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "INSERT banned (banner, banned) VALUES (?,?)", banned.ID)

}