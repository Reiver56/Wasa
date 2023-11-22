package database


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
		followers TEXT,
		following TEXT,
		photos TEXT
		





