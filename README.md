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

Performance testing:
1. Download wrk image: `docker pull williamyeh/wrk`
2. Generate drivers: `curl http://localhost:8080/testapi/drivers -d '{"count": 1000, "cleanup": true}'`
3. Generate requests: `for((i=1;i<=1000;i++)); do echo "/driver/$i"; done > perftest/paths.txt`
4. Run test:
```
docker run --rm \
    --net=host \
    -v `pwd`/perftest:/ego \
    williamyeh/wrk -t4 -c100 -d10s -s /ego/script.lua http://localhost:8080/`
```
