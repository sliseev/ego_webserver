FROM golang:1.20.0 AS builder

WORKDIR /usr/src/ego
COPY . .
RUN go mod tidy
RUN make build

FROM ubuntu AS runtime

WORKDIR /opt/app/ego
COPY --from=builder /usr/src/ego/ego_server .

ENTRYPOINT [ "./ego_server" ]
