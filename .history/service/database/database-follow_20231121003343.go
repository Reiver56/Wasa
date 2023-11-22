package database


//  this fuction define get follow of a user
func (db *_database) GetFollow(id_user string) ([]Profile, error) {
	 rows, err := db.c.Query("SELECT follower FROM follow WHERE followed = ?", id_user)
}