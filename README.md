# Simple DB

This is a pet project for exercising my skills in Golang while trying to implement an in-memory database, client and protocol.

![Show me what you got](./ricky.gif)

Currently there are 2 engines available:

## Key Value Engine

Example:

```
$ go run cmd/server/main.go -port 3000 -engine key-value
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

## Rolling Window Engine

Starts a server with rolling window storage provided through [Horus](https://github.com/ricardohsd/horus).

This engine provides the following commands:
```
RWSET key-name SECONDS
RWADD key-name VALUE-FLOAT
RWAVG key-name
RWDEL key-name
```

Example:

Starting the server.
```
$ go run cmd/server/main.go -engine rolling-window
2018/08/20 21:43:57 Starting server on address 127.0.0.1:3000, with engine rolling-window
2018/08/20 21:44:10 Processing Set transactions
2018/08/20 21:44:19 Processing Add transactions
2018/08/20 21:44:24 Processing Average transactions
2018/08/20 21:44:29 Processing Average transactions
2018/08/20 21:44:42 Processing Average 50.0
2018/08/20 21:44:48 Processing Average transactions
2018/08/20 21:45:03 Processing Add transactions
2018/08/20 21:45:08 Processing Average transactions
2018/08/20 21:45:22 Processing Add transactions
2018/08/20 21:45:25 Processing Average transactions
2018/08/20 21:45:37 Processing Average transactions
2018/08/20 21:45:52 Processing Add transactions
2018/08/20 21:45:55 Processing Average transactions
```

Connecting with the client.
```
λ go run cmd/client/main.go
> rwset transactions 60
 OK
> rwadd transactions 10.10
 OK
> rwavg transactions
 0.17
> rwavg transactions
 0.17
> rwavg 50.0
 key not found
> rwavg transactions 50.0
 0.17
> rwadd transactions 50.0
 OK
> rwavg transactions
 1.00
> rwadd transactions 5.0
 OK
> rwavg transactions
 0.92
> rwavg transactions
 0.92
> rwadd transactions 30.0
 OK
> rwavg transactions
 1.42
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