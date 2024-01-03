package database

import "fmt"

func (db *appdbimpl) GetListPhoto(id_user string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT * FROM photo WHERE user_id = ? ORDER BY date DESC", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// read all rows with user result
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
			fmt.
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

// create a new photo in database
func (db *appdbimpl) UploadPhoto(p Photo) (int64, error) {

	res, err := db.c.Exec(`INSERT INTO photo(user_id, date) VALUES(?, ?)`, p.User_ID, p.Date)
	if err != nil {
		return -1, err
	}
	photo_id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return photo_id, nil

}

// delete a photo in database
func (db *appdbimpl) DeletePhoto(user_id string, id_photo int64) error {

	_, err := db.c.Exec(`DELETE FROM photo WHERE id_photo = ?`, id_photo)
	if err != nil {
		return err
	}
	return nil
}

// get a photo in database
func (db *appdbimpl) GetPhoto(tphoto_id int64, id_req string) (Photo, error) {
	var photo Photo
	err := db.c.QueryRow(`SELECT * FROM photo WHERE id_photo = ? AND user_id NOT IN (SELECT banner FROM ban WHERE banned = ? ) `, tphoto_id, id_req).Scan(&photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}
