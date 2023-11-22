package database


//  this fuction define get follow of a user
func (db *appdbimpl) GetFollow(id_user string) ([]Profile, error) {
	 rows, err := db.c.Query("SELECT follower FROM follow WHERE followed = ?", id_user)
	 if err != nil {
		 return nil, err
	 }
	 
}