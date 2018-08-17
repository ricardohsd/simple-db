# Simple DB

This is a pet project for exercising my skills in Golang while trying to implement an in-memory database, client and protocol.

![Show me what you got](./ricky.gif)

Example:

```
$ go run cmd/server/main.go -port 3000
2018/08/15 22:42:04 Starting server on address :3000
2018/08/15 22:42:16 Processing SET name, john
2018/08/15 22:42:20 Processing GET name
2018/08/15 22:42:23 Processing GET age
2018/08/15 22:42:26 Processing DEL name
2018/08/15 22:42:30 Processing GET name
2018/08/15 22:42:33 Processing DEL name
2018/08/15 22:43:02 EOF. Closing connection
```

```
$ go run cmd/client/main.go -port 3000
> SET name john
 OK
> GET name
 john
> GET age
 key not found
> DEL name
 OK
> GET name
 key not found
> DEL name
 key not found
```

# Debug Proxy

A minimalistic debugging proxy can be started by:

```
$ go run cmd/proxy/main.go -listen 3001 -backend 300
2018/08/17 19:25:17 Started proxy on address :3001
2018/08/17 19:25:21 Statistics map[]
2018/08/17 19:25:26 Forwarding message john to client
2018/08/17 19:25:26 Statistics map[GET:1]
2018/08/17 19:25:30 Forwarding message OK to client
2018/08/17 19:25:30 Statistics map[SET:1 GET:1]
2018/08/17 19:25:33 Forwarding message 30 to client
2018/08/17 19:25:33 Statistics map[GET:2 SET:1]

```

The proxy will count each message being forwarded from the client to the server.