package database

//  Database fuction that add a new user in dat
func (db *appdbimpl) CreateUser(username string, password string) error{
	_, err := db.c.Exec("INSERT INTO users (username,password) VALUES (?,?), username, password")

	if err != nil {
		return err
	}
	return nil
}