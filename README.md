# SYNOPSIS

go http logging - some opinion how to log http requests/responses

This isn't a library or anything fancy, just a mini tutorial on getting at the http
response code...

# DESCRIPTION

Logging the status code for an HTTP response in Go turns out to be a little
tricker than first anticipated. You don't get access to the `Response` object
after your handler, rather you only have the `ResponseWriter` which is like a
byte stream.

So I ripped out the `responseLogger` from Gorilla's
[handlers](https://github.com/gorilla/handlers) package and made it so that I
can get the response status code and the response size after our handler runs.

Also I went with a structured logger
[logrus](https://github.com/sirupsen/logrus) as that's closer to what I'm used
to using in node land, e.g., [bunyan](https://github.com/trentm/node-bunyan).

# EXAMPLE

## Server
```sh
± go run main.go response_logger.go
{"level":"info","msg":"listening on port 9990","time":"2016-12-10T13:20:35-05:00"}
{"addr":"[::1]:49290","code":200,"level":"info","method":"GET","msg":"HTTP","size":14,"time":"2016-12-10T13:20:38-05:00","took":"24.466µs","url":"/"}
{"addr":"[::1]:49297","code":404,"level":"info","method":"GET","msg":"HTTP","size":15,"time":"2016-12-10T13:20:40-05:00","took":"9.058µs","url":"/status"}
```

## Client

```sh
± curl -v -H'Accept: application/json' localhost:9990/
*   Trying ::1...
* Connected to localhost (::1) port 9990 (#0)
> GET / HTTP/1.1
> Host: localhost:9990
> User-Agent: curl/7.49.1
> Accept: application/json
>
< HTTP/1.1 200 OK
< Date: Sat, 10 Dec 2016 18:21:22 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
Up and running%
```
