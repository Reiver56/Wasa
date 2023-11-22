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
		id
);`

//  ---------Comment Table----------------

var CommentTable = `CREATE TABLE IF NOT EXISTS comments
(
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		photo_id TEXT,
		comment TEXT
);`

//  ---------Follow Table----------------

var FollowTable = `CREATE TABLE IF NOT EXISTS follow
(
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		following TEXT
);`

//  ---------Ban Table----------------

var BanTable = `CREATE TABLE IF NOT EXISTS ban
(
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		banned TEXT
);`
