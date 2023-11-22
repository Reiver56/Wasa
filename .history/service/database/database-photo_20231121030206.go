package database

func (db *appdbimpl) GetListPhoto(reqUser string, id_user string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photo WHERE ", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
	var photos []Photo
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.Id_photo, &photo.Nickname, &photo.User_ID, &photo.Date); err != nil {
			return nil, err
		}
		photo.Likes, _ = db.GetLikes(photo.Id_photo)
		photo.Comments, _ = db.GetComments(photo.Id_photo)
		photos = append(photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}