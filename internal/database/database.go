package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(host, user, password, dbname string, port int, verbose bool) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		host,
		user,
		password,
		dbname,
		port,
	)

	log_level := logger.Silent
	if verbose {
		log_level = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(log_level),
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Driver{}); err != nil {
		return nil, err
	}

	return db, nil
}
