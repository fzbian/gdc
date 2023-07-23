package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_DATABASE)

	// Open a connection to the database using the connection string.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Desactiva el log de GORM
	})
	if err != nil {
		// If there was an error opening the connection, panic.
		panic(err)
	}

	// Return the database connection.
	return db
}
