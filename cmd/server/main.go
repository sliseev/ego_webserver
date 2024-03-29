package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/vl80s/ego_webserver/internal/database"
	"github.com/vl80s/ego_webserver/internal/server"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Commit string

	host       = flag.String("host", "0.0.0.0", "address to listen on")
	port       = flag.Int("port", 8080, "port to listen on")
	verbose    = flag.Bool("verbose", false, "enable extended logging")
	standalone = flag.Bool("standalone", false, "run without dependencies")
)

func NewLogger() (*zap.Logger, error) {
	if *verbose {
		return zap.NewDevelopment()
	} else {
		return zap.NewProduction()
	}
}

func NewDatabase(logger *zap.Logger) (*gorm.DB, error) {
	if *standalone {
		logger.Info("Database in-memory enabled (standalone)")
		return database.LtConnect(*verbose)
	}

	logger.Info("Database Postgres used (env)",
		zap.String("DB_HOST", os.Getenv("DB_HOST")),
		zap.String("DB_PORT", os.Getenv("DB_PORT")),
	)

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 5432 // default Postgres port
	}
	return database.PgConnect(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port,
		*verbose)
}

// @title		EGO Service
// @version		1.0
// @description	This is a simple REST server for upping up skills.
// @contact.name	Sergey Liseev
// @contact.email	sergey_liseev@epam.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host		localhost:8080
func main() {
	flag.Parse()

	logger, err := NewLogger()
	if err != nil {
		fmt.Printf("FATAL: Can't initialize logger: %v", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Start service", zap.String("commit", Commit))

	db, err := NewDatabase(logger)
	if err != nil {
		logger.Fatal("Connection to database failed", zap.Error(err))
	}

	srv := server.New(db, logger)
	logger.Info("Running service at", zap.String("host", *host), zap.Int("port", *port))
	logger.Fatal("Stop service", zap.Error(srv.Run(*host, *port)))
}
