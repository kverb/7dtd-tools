package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kverb/7dtd-server-tools/parser"
)

type InfoHandler struct{}

var last = []byte{}

func (hh InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	host := query.Get("host")
	port := query.Get("port")
	filter := query.Get("filter")
	hostPort := host + ":" + port
	resp, err := parser.QueryServerBytes(hostPort, filter)
	if err != nil {
		resp = last
	}
	w.Write(resp)
}

func main() {
	arguments := os.Args
	var listenPort string
	if len(arguments) == 1 {
		listenPort = "8787"
	} else {
		listenPort = arguments[1]
	}

	s := http.Server{
		Addr:         ":" + listenPort,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      InfoHandler{},
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
