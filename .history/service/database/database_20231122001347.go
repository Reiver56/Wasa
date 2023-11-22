/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// Return true if user exists, false otherwise
	ExistUser(username string) (bool, error)

	// Create a new User in the DB
	CreateUser(user User) error

	// Check if user exists
	CheckUser(user User) bool

	// Get a user (id)
	GetUser(username string) (User, error)

	// Get a user (nickname)
	GetNickname(id string) (string, error)

	// Modify user's username
	SetNewUsername(userID string, username string) error
	
	// Get user's stream
	GetMyStream(user User) ([]Photo, error)
	
	// Get photo
	GetPhoto(tphoto_id int64, id_req string) (Photo, error)
	
	// Put a new photo in datbase
	UploadPhoto(p Photo) (int64, error)

	// Delete a photo in database
	DeletePhoto(user_id string, id_photo int64) error
	
	// Get follow of a user
	GetFollow(id_user string) ([]User, error)

	// Get user's following another user
	GetFollowing(id_user string) ([]User, error)

	
	
	// User follow another user
	Follow(id_user string, id_followed string) error
	
	// User unfollow another user
	Unfollow(id_user string, id_followed string) error
	
	// Get a list of likes of a photo
	GetListLikes(id_photo int) ([]Like, error)

	// Get a list of a photo
	GetListPhoto(id_user string) ([]Photo, error)

	// Get a list of users that match with the search
	GetUserProfile(u1 User, u2 User) ([]Profile, error)
	
	// Get a list of comments of a photo
	GetListComments(id_photo int64) ([]Comment, error)

	// Put a new comment in database
	PostComment(id_photo int64, id_user string, c Comment) (int64, error)

	// Delete a comment in database
	DeleteComment(id_photo int64, id_user string, id_comment int64) (error)
	
	// Put a like
	LikePhoto(photo_id int, user User) error
	
	// Delete a like
	UnlikePhoto(like Like) error

	// Verify if a user is banned
	IsBanned(id_banned string, id_banner string ) (error,bool)
	
	// Ban a user
	BanUser(banner User, banned User) error

	// Unban a user
	UnbanUser(banner User, banned User) error

	//Ping the database to check if is alive
	Ping() error
}
type appdbimpl struct {
	c   *sql.DB
	ctx context.Context
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if the table exists
	var tableName uint8
	err := db.QueryRow(`SELECT count(name) FROM sqlite_master WHERE type='table';`).Scan(&tableName)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	if tableName != 6 {
		// Create the user table
		_, err := db.Exec(UserTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		// Create the photo table
		_, err = db.Exec(PhotoTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// Create the like table
		_, err = db.Exec(LikeTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// Create the comment table
		_, err = db.Exec(CommentTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// Create the follow table
		_, err = db.Exec(FollowTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// Create the ban table
		_, err = db.Exec(BanTable)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
