package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}
	return nil
}


func GetSqLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// initialize logger 
	logger = NewLogger(p)
	return logger
}