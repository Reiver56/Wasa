package database


//---------User Table----------------

var UserTable = `CREATE TABLE IF NOT EXISTS users 
(
		id SERIAL PRIMARY KEY,
		nickname TEXT 
);`

//---------Profile Table--------------

var ProfileTable = `CREATE TABLE IF NOT EXISTS profile
(
		id SERIAL PRIMARY KEY,
		nickname TEXT,
		followers TEXT,
		following TEXT,
		photos TEXT
);`

//---------Photo Table----------------

var PhotoTable = `CREATE TABLE IF NOT EXISTS photo
(
		id SERIAL PRIMARY KEY,
		





