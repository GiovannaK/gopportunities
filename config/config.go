package config

import "gorm.io/gorm"

var (
	// DB is the variable that holds the database connection
	db *gorm.DB
)

func Init() error {
	return nil
}