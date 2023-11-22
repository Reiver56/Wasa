package database

func (db *appdbimpl) GetListPhoto(reqUser string, id_user string) ([]Photo, error) {
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
		comments, err := db.GetComments(photo.Id_photo)
		if err != nil {
			return nil, err
		}
		likes, err := db.GetListLikes(photo.Id_photo)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return photos, nil
}