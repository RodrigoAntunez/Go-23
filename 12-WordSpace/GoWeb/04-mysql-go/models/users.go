package models

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

// Esquema en SQL
const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSEIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	reg_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)`
