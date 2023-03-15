package main

import (
	"log"
	"os"
	"strconv"

	"github.com/vl80s/ego_webserver/internal/database"
	"github.com/vl80s/ego_webserver/internal/server"
)

var Commit string = "<none>"

func main() {
	log.Printf("Load server [commit %s]", Commit)

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := database.Connect(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		port)
	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}

	srv := server.New(db)
	srv.Run("0.0.0.0", 8081)
}
