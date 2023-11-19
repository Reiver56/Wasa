package database

//create a like in database
func (db *appdbimpl) likePhoto(like Like) error {
	
	_, err := db.c.Exec(`INSERT INTO likes(id_photo, id_user) VALUES(?, ?)`,)
	if err != nil {
		return err
	}
	return nil

}

