package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening on port %v\n", listener.Addr())
	}

	server := &http.Server{
		Handler:        Router{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

type Router struct{}

type MyDataResponse struct {
	Value string `json:"value,omitempty"`
}

func (Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	switch r.URL.Path {
	case "/sample1":
		value, err := json.Marshal(MyDataResponse{"This is sample 1"})
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(value)
		}
	case "/sample2":
		value, err := json.Marshal(MyDataResponse{"This is sample 2"})
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(value)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		_, err = io.WriteString(w, `{"error": "Path not found"}`)
	}

	if err != nil {
		log.Println(err)
	}
}
