# SYNOPSIS

go http logging - some opinion how to log http requests/responses

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
