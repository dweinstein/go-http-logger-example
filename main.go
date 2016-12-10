package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)

	r := mux.NewRouter()

	r.Handle("/", StatusHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotImplemented)
	log.Println("listening on port 9990")
	log.Fatal(http.ListenAndServe(":9990", HTTPLogger(r)))
}

func HTTPLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tic := time.Now()
		var logger loggingResponseWriter = &responseLogger{w: w}
		handler.ServeHTTP(logger, r)
		log.WithFields(log.Fields{
			"code":   logger.Status(),
			"took":   fmt.Sprintf("%v", time.Since(tic)),
			"addr":   r.RemoteAddr,
			"method": r.Method,
			"url":    fmt.Sprintf("%s", r.URL),
			"size":   logger.Size(),
		}).Info("HTTP")
	})
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Up and running"))
})
