# Golang with Consul to achieve mTLS

*mTLS - service-to-service TLS encryption

`server` - directory with http server\
`client` - directory with the client that calls server via consul

## 1. Run consul
```
consul agent \                                                                                                                                                                                                                                          1 ✘  2h 40m 58s
-server \
-bootstrap \
-ui \
-config-file consul.config.json \
--data-dir /tmp/consul \
-client 0.0.0.0
```

## 2. Run server
```
go run server/server.go
```

## 3. Run client
```
go run client/client.go
```
Client will return 404, which means success, since the handshake was completed successfully