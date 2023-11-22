package database

func (db *appdbimpl) GetListComments(ID_comment int) ([]Comment, error) {
	rows, err := db.c.Query("SELECT id_user FROM likes WHERE id_photo = ?", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
	var likes []Like
	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.ID_like, &like.ID_user); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return likes, nil
}