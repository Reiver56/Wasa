package database

// this fuction define get follow of a user
func (db *appdbimpl) GetFollow(id_user string) ([]User, error) {
	rows, err := db.c.Query("SELECT id, nickname FROM users WHERE id IN (SELECT follower FROM follow WHERE followed = ?)", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// read all rows with user result
	var followers []User

	for rows.Next() {
		var follower User
		if err := rows.Scan(&follower.ID, &follower.Nickname); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}

// this fuction define get user's following another user
func (db *appdbimpl) GetFollowing(id_user string) ([]User, error) {
	rows, err := db.c.Query("SELECT id, nickname FROM users WHERE id IN (SELECT followed FROM follow WHERE follower = ?)", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// read all rows with user result
	var following []User
	for rows.Next() {
		var followed User
		if err := rows.Scan(&followed.ID, &following); err != nil {
			return nil, err
		}
		following = append(following, followed)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return following, nil
}

// function that define user follow another user
func (db *appdbimpl) Follow(id_user string, id_followed string) error {
	_, err := db.c.Exec("INSERT INTO follow(follower, followed) VALUES(?, ?)", id_user, id_followed)
	if err != nil {
		return err
	}
	return nil
}

// function that define user unfollow another user
func (db *appdbimpl) Unfollow(id_user string, id_followed string) error {
	_, err := db.c.Exec("DELETE FROM follow WHERE follower = ? AND followed = ?", id_user, id_followed)
	if err != nil {
		return err
	}
	return nil
}
