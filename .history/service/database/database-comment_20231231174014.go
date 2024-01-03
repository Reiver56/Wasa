package database

func (db *appdbimpl) GetListComments(id_photo int64) ([]Comment, error) {
	rows, err := db.c.Query("SELECT *FROM comments WHERE id_photo = ?", id_photo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// read all rows with user result
	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ID_comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

// create a new comment in database
func (db *appdbimpl) PostComment(id_photo int64, id_user string, c Comment) (int64, error) {

	res, err := db.c.Exec(`INSERT INTO comments(id_photo, id_user, text) VALUES(?, ?, ?)`, id_photo, id_user, c.Text_comment)
	if err != nil {
		return -1, err
	}
	comment_id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return comment_id, nil

}

func (db *appdbimpl) DeleteComment(id_photo int64, id_user string, id_comment int64) error {

	_, err := db.c.Exec(`DELETE FROM comments WHERE (id_comment = ?, id_photo = ?, id_user = ?)`, id_photo, id_user, id_comment)
	if err != nil {
		return err
	}
	return nil

}
