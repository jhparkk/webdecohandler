package main

import (
	"log"
	"net/http"
	"time"

	"jhpark.sinsinway.com/webdecohandler/decohandler"
	"jhpark.sinsinway.com/webdecohandler/myapp"
)

func logger1(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Start")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] End time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Start")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] End time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := decohandler.NewDecoHandler(mux, logger1)
	h = decohandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
