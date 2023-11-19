package database

//return a list of matching users with u2 (that is the user that is searching)
func (db *appdbimpl) getUserProfile(u1 User, u2 User) ([]Profile, error) {
	rows, err := db.c.Query(`SELECT * FROM users WHERE (nickname LIKE ? OR id != ?) AND id NOT IN (SELECT banner FROM banned WHERE banned = ?)`, u2.Nickname+"%", u1.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var profiles []Profile
	for rows.Next() {
		var profile Profile
		err := rows.Scan(&profile.ID, &profile.Nickname)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return profiles, nil
	
	
}