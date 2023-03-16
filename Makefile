build:
	go build -o ego_server -ldflags="-X main.Commit=$(shell git rev-parse HEAD)" cmd/server/main.go

swagger:
	swag init -d cmd/server/,internal/server/ -g main.go --ot yaml
