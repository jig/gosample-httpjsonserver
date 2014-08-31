package main

import (
	"fmt"
	"github.com/gocraft/web"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Listening on port %v\n", listener.Addr())
	}

	router := web.New(Context{}). // Create your router
					Middleware(web.LoggerMiddleware).     // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get("/", (*Context).SayHello)         // Add a route

	server := &http.Server{
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
