package database



func (db *appdbimpl) ExistUser(username string, password string) (bool, error){
	var existUser string
	err := db.c.QueryRow(query_FINDUSER, username, password).Scan(&existUser)
	if err != nil && errors.Is(err,sql.ErrNoRows){
		return false, err
		}
	return existUser != "",err	 
}

