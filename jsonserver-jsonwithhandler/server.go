package main

import (
	"encoding/json"
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

	sampleMux := http.NewServeMux()
	sampleMux.HandleFunc("/sample/1", func(w http.ResponseWriter, req *http.Request) {
		value, err := json.Marshal(MyDataResponse{"This is sample 1"})
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(value)
		}
	})

	sampleMux.HandleFunc("/sample/2", func(w http.ResponseWriter, req *http.Request) {
		value, err := json.Marshal(MyDataResponse{"This is sample 2"})
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(value)
		}
	})

	serverMux := http.NewServeMux()
	serverMux.Handle("/sample/", sampleMux)
	serverMux.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("/tmp"))))

	server := &http.Server{
		Handler:        serverMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

type MyDataResponse struct {
	Value string `json:"value,omitempty"`
}
