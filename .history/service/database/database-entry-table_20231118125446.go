package database

/*

 This file contains the SQL queries to create the tables for the application's database.
  The tables include users, profile, photo,like,comment,follow,and ban table, each with their respective columns.
 The queries are written in SQL and stored as string variables.

*/

//  ---------User Table----------------

var UserTable = `CREATE TABLE IF NOT EXISTS users 
(
		id VARCHAR(16) NOT NULL PRIMARY KEY,
		nickname VARCHAR(16) NOT NULL 
);`


//  ---------Photo Table----------------

var PhotoTable = `CREATE TABLE IF NOT EXISTS photo
(
		id_photo INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id VARCHAR(16) NOT NULL,
		date DATETIME NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`

//  ---------Like Table----------------

var LikeTable = `CREATE TABLE IF NOT EXISTS likes
(
		id_photo INTEGER NOT NULL,
		id_user VARCHAR(16) NOT NULL,
		PRIMARY KEY (id_photo, id_user),
		FOREIGN KEY (id_photo) REFERENCES photo(id_photo) ON DELETE CASCADE,
		FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE
);`

//  ---------Comment Table----------------

var CommentTable = `CREATE TABLE IF NOT EXISTS comments
(
		id_comment INTEGER PRIMARY KEY AUTOINCREMENT,
		id_photo INTEGER NOT NULL,
		id_user VARCHAR(16) NOT NULL,
		text VARCHAR(255) NOT NULL,
		FOREIGN KEY (id_photo) REFERENCES photo(id_photo) ON DELETE CASCADE,
		FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE
);`

//  ---------Follow Table----------------

var FollowTable = `CREATE TABLE IF NOT EXISTS follow
(
		follower VARCHAR(16) NOT NULL,
		followed VARCHAR(16) NOT NULL,
		PRIMARY KEY (follower, followed),
		FOREIGN KEY (follower) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (followed) REFERENCES users(id) ON DELETE CASCADE
);`

//  ---------Ban Table----------------

var BanTable = `CREATE TABLE IF NOT EXISTS ban
(
		banned VARCHAR(16) NOT NULL,
		
);`
