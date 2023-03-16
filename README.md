EGO WebServer
=============

Skill up project. Uses the following staff:
- gin-gonic web server
- gorm database ORM
- swag docs generator

Prepare .env file based on template:

```
export DB_HOST=
export DB_PORT=
export DB_USER=
export DB_PASSWORD=
export DB_NAME=
```

To generate swagger you must install the tool:
```
go install github.com/swaggo/swag/cmd/swag@latest
```
