build:
	go build -ldflags="-X main.Commit=$(shell git rev-parse HEAD)" cmd/server/main.go