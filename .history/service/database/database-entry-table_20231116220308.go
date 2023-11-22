
package database


 This file contains the SQL queries to create the tables for the application's database.
//   The tables include users, profile, and photo, each with their respective columns.
//   The queries are written in SQL and stored as string variables.


//  ---------User Table----------------

var UserTable = `CREATE TABLE IF NOT EXISTS users 
(
		id SERIAL PRIMARY KEY,
		nickname TEXT 
);`

//  ---------Profile Table--------------

var ProfileTable = `CREATE TABLE IF NOT EXISTS profile
(
		id SERIAL PRIMARY KEY,
		nickname TEXT,
		followers TEXT,
		following TEXT,
		photos TEXT
);`

//  ---------Photo Table----------------

var PhotoTable = `CREATE TABLE IF NOT EXISTS photo
(
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		likes TEXT,
		comments TEXT,
		date TEXT
);`






