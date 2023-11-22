package database


//  this fuction define get follow of a user
func (db *_database) GetFollow(id_user string) ([]Profile, error) {
	 rows, err := bd.c.Q