package database

func (db *appdbimpl) GetListPhoto(id_user string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photo WHERE user_id = ? ORDER BY date DESC", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
	var photos []Photo
	for rows.Next() {
		var photo Photo
		
		if err := rows.Scan(&photo.Id_photo, &photo.User_ID, &photo.Date); err != nil {
			return nil, err
		}
		comments, err := db.GetListComments(photo.Id_photo)
		if err != nil {
			return nil, err
		}
		likes, err := db.GetListLikes(photo.Id_photo)
		if err != nil {
			return nil, err
		}
		photo.Comments = comments
		photo.Likes = likes
		photos = append(photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}

//  create a like in database
func (db *appdbimpl) UploadPhoto(photo_id int, user User) error {
	
	_, err := db.c.Exec(`INSERT INTO likes(id_photo, id_user) VALUES(?, ?)`, photo_id, user.ID)
	if err != nil {
		return err
	}
	return nil

}