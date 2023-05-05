package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PgConnect(host, user, password, dbname string, port int, verbose bool) (*gorm.DB, error) {
	return connect(verbose, func() gorm.Dialector {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d",
			host,
			user,
			password,
			dbname,
			port,
		)
		return postgres.Open(dsn)
	})
}

func LtConnect(verbose bool) (*gorm.DB, error) {
	return connect(verbose, func() gorm.Dialector {
		return sqlite.Open("file::memory:?cache=shared")
	})
}

func connect(verbose bool, dialect func() gorm.Dialector) (*gorm.DB, error) {
	log_level := logger.Silent
	if verbose {
		log_level = logger.Info
	}

	db, err := gorm.Open(dialect(), &gorm.Config{
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
