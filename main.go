package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kverb/7dtd-server-tools/parser"
)

type InfoHandler struct{}

func (hh InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	host := query.Get("host")
	port := query.Get("port")
	filter := query.Get("filter")
	hostPort := host + ":" + port
	respMap, err := parser.QueryServer(hostPort)
	if err != nil {
		respMap = map[string]string{"error": err.Error()}
	}

	if len(filter) > 0 {
		v, ok := respMap[filter]
		if ok {
			respMap = map[string]string{filter: v}
		}
	}
	b, err := json.MarshalIndent(respMap, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
	w.Write([]byte(b))
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
