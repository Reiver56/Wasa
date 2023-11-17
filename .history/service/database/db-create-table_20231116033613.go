package database

var sql_TABLEUSER = `CREATE TABLE IF NOT EXISTS User
(
	
	username STRING NOT NULL UNIQUE,
	PRIMARY KEY(userID)
);`