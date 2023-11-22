package database

func (db *appdbimpl) GetListComments(id_photo int64) ([]Comment, error) {
	rows, err := db.c.Query("SELECT id_comment FROM comments WHERE id_photo = ?", id_photo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//read all rows with user result
	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID_comment, &comment.ID_user); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

//  create a new comment in database
func (db *appdbimpl) PutComment(id_photo int64, id_user string, text string) (int64, error) {
	
	res, err := db.c.Exec(`INSERT INTO comments(id_photo, id_user, comment) VALUES(?, ?, ?)`, id_photo, id_user, text)
	if err != nil {
		return -1,err
	}
	comment_id, err := res.LastInsertId()
	
	if err != nil {
		return -1,err
	}

	return comment_id, nil

}