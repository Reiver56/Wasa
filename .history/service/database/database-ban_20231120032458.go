package database

func (db *appbimp1) BanUser(banner User, banned User) error {
	_, err := db.c.ExecContext(db.ctx, "UPDATE banned SET banned = 1 WHERE id = ?", banned.ID)
}