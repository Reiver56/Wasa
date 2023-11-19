package database

//create a like in database
func (db *appdbimpl) likePhoto(like Like) error {
	
	_, err := db.c.Exec(`INSERT INTO likes(id_photo, id_user) VALUES(?, ?)`, like.ID_like, like.ID_user)
	if err != nil {
		return err
	}
	return nil

}
//delete a like in database
func (db *appdbimpl) unlikePhoto(like Like) error {
	

