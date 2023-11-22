package database

//  create a like in database
func (db *appdbimpl) LikePhoto(photo_id int, user User) error {
	
	_, err := db.c.Exec(`INSERT INTO likes(id_photo, id_user) VALUES(?, ?)`, like.ID_like, like.ID_user)
	if err != nil {
		return err
	}
	return nil

}
//  delete a like in database
func (db *appdbimpl) UnlikePhoto(like Like) error {
	
	_, err := db.c.Exec(`DELETE FROM likes WHERE id_photo = ? AND id_user = ?`, like.ID_like, like.ID_user)
	if err != nil {
		return err
	}
	return nil

}

