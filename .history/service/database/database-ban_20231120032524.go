package database

func (db *appbimp1) BanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "INSERT banned (banner, banned) banned = 1 WHERE id = ?", banned.ID)
}