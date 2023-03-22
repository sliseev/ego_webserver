EGO WebServer
=============

Skills up project. Uses the following staff:
- docker compose
- gin-gonic web server
- gorm database ORM
- swag docs generator
- prometheus + grafana
- wrk performance testing tool
- nginx reverse proxy with ssl

### Prepare .env file based on template (see sample values):

```
export DB_HOST=127.0.0.1 # for local run only, not used by docker
export DB_PORT=5432      # for local run only, not used by docker

export DB_USER=anybody
export DB_PASSWORD=qwerty
export DB_NAME=ego
```

### Prepare self-signed certificate for nginx:

```
$ mkdir .cert
# create certificate using any instruction from the internet,
# e.g. https://devopscube.com/create-self-signed-certificates-openssl/
# finally you need to have two files in ./.cert: server.crt & server.key
```

### Start service locally:

```
$ source .env
$ make build
$ ./ego_server
```

### Start service in docker:

```
$ source .env
$ docker compose up
```

### Request sample:

```
$ curl http://localhost:8080/driver/count
$ curl -k https://localhost:8443/driver/count
```

### Swagger generation:

```
$ go install github.com/swaggo/swag/cmd/swag@latest
$ make swagger
```

### Grafana usage:

1. Login: admin/admin
2. Create datasource: Type: Prometheus, URL: http://prometheus:9090
3. Import dashboard: grafana.json

### Performance testing:

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
