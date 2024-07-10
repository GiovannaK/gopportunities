package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}


func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// initialize logger 
	logger = NewLogger(p)
	return logger
}