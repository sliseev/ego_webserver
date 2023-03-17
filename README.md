EGO WebServer
=============

Skill up project. Uses the following staff:
- gin-gonic web server
- gorm database ORM
- swag docs generator
- prometheus + grafana

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

Grafana:
1. Login: admin/admin
2. Create datasource: Type: Prometheus, URL: http://prometheus:9090
3. Import dashboard: grafana.json
