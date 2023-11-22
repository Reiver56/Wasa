package database

func (db *appdbimpl) GetListComments(id_photo int64) ([]Comment, error) {
	rows, err := db.c.Query("SELECT id_comment FROM comments WHERE id_photo = ?", id_comment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
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