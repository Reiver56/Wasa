package database

func (db *appdbimpl) GetListLikes(id_photo int) ([]Like, error) {
	rows, err := db.c.Query("SELECT id_user FROM likes WHERE id_photo = ?", id_photo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
	var likes []Like
	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.ID_photo, &like.ID_user); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return likes, nil
}
//  create a like in database
func (db *appdbimpl) LikePhoto(photo_id int, user User) error {
	
	_, err := db.c.Exec(`INSERT INTO likes(id_photo, id_user) VALUES(?, ?)`, photo_id, user.ID)
	if err != nil {
		return err
	}
	return nil

}
//  delete a like in database
func (db *appdbimpl) UnlikePhoto(like Like) error {
	
	_, err := db.c.Exec(`DELETE FROM likes WHERE id_photo = ? AND id_user = ?`, like.ID_photo, like.ID_user)
	if err != nil {
		return err
	}
	return nil

}
func (db *appdbimpl) LikeCheck(id string) (bool, error) {
	var existUser string
	err := db.c.QueryRow(`SELECT id_photo FROM likes WHERE id_ = ?`, nickname).Scan(&existUser)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return existUser != "", err

}

