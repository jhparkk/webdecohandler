package main

import (
	"log"
	"net/http"
	"time"

	"jhpark.sinsinway.com/webdecohandler/decohandler"
	"jhpark.sinsinway.com/webdecohandler/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Start")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] End time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := decohandler.NewDecoHandler(mux, logger)
	return h
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
