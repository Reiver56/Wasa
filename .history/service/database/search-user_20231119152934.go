package database

import "fmt"

//  return a list of matching users with u2 (that is the user that is searching)
func (db *appdbimpl) GetUserProfile(u1 User, u2 User) ([]Profile, error) {
	
	
	rows, err := db.c.Query(`SELECT * FROM users WHERE ((id LIKE ?) OR (nickname LIKE ?)) AND id NOT IN (SELECT banner FROM banned WHERE banned = ?)`, u2.ID+"%", u2.ID+"%", u1.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)
	var profiles []Profile
	for rows.Next() {
		var profile Profile
		err := rows.Scan(&profile.ID, &profile.Nickname)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
		fmt.Println(profiles)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return profiles, nil
	
	
}