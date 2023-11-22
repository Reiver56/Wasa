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

	CreateUser(u User) error

	ExistUser(nickname string) (bool, error)

	SetNewUsername(userID string, username string) error

	GetUser(username string) (User, error)

	GetNickname(id string) (string, error)

	CheckUser(user User) (bool)

	GetMyStream(user User) ([]Photo, error)

	GetUserProfile(u1 User, u2 User) ([]Profile, error)

	GetListPhoto(id_user string) ([]Photo, error)

	UploadPhoto(p Photo) (int64, error)

	DeletePhoto(user_id string, id_photo int64) error
	
	GetPhoto(tphoto_id int64, id_req string) (Photo, error)

	GetListLikes(id_photo int64) ([]Like, error)

	LikePhoto(photo_id int64, user User) error

	UnlikePhoto(like Like) error

	GetFollow(id_user string) ([]User, error)

	GetFollowing(id_user string) ([]User, error)

	Follow(id_user string, id_followed string) error

	Unfollow(id_user string, id_followed string) error

	GetListComments(id_photo int64) ([]Comment, error)

	PostComment(id_photo int64, id_user string, c Comment) (int64, error)

	
	// Ping the database to check if is alive
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
