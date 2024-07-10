package config

import (
	"os"

	"github.com/GiovannaK/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// check if the database already exists

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist")
		// create the database file 
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Error creating the database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Error creating the database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	// create a new connection to the database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error connecting to the database: %v", err)
		return nil, err
	}

	// migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error migrating the schema: %v", err)
		return nil, err
	}

	logger.Info("Connected to the database")

	// return the connection
	return db, nil
}
