package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// Db main connection
	Db = Connect()
)

// Connect opens a connection to the database.
// The function returns a *gorm.DB object.
func Connect() *gorm.DB {

	DB_USER := "root"
	DB_PASSWORD := ""
	DB_HOST := "localhost"
	DB_PORT := "3306"
	DB_DATABASE := " julio"

	// Create a connection string.
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_DATABASE)

	// Open a connection to the database using the connection string.
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		// If there was an error opening the connection, panic.
		panic(err)
	}

	// Return the database connection.
	return db
}
