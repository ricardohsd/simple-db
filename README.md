# Simple DB

This is a pet project for exercising my skills in Golang while trying to implement an in-memory database, client and protocol.

![Show me what you got](./ricky.gif)

Example:

```
$ go run cmd/server/main.go
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
$ go run cmd/client/main.go
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