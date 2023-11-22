package database


//  this fuction define get follow of a user
func (db *appdbimpl) GetFollow(id_user string) ([]Profile, error) {
	 rows, err := db.c.Query("SELECT follower FROM follow WHERE followed = ?", id_user)
	 if err != nil {
		 return nil, err
	 }
	 defer rows.Close()
	 //  read all rows with user result
	 var followers []Profile
	 for rows.Next() {
		 var follower Profile
		 if err := rows.Scan(&follower.ID); err != nil {
			 return nil, err
		 }
		 followers = append(followers, follower)
	 }
	 if err := rows.Err(); err != nil {
		 return nil, err
	 }
	 return followers, nil
}

//  this fuction define get user's following another user
func (db *appdbimpl) GetFollowing(id_user string) ([]Profile, error) {
	rows, err := db.c.Query("SELECT followed FROM follow WHERE follower = ?", id_user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//  read all rows with user result
	var following []Profile
	for rows.Next() {
		var followed Profile
		if err := rows.Scan(&followed.ID); err != nil {
			return nil, err
		}
		following = append(following, followed)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return following, nil
}
//  function that define user 
