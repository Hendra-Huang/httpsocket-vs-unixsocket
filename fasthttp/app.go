package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

type Server struct {
	Type string
}

func (s Server) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	body := "Hello World " + s.Type + "\n"
	fmt.Fprint(ctx, body)
}

func main() {
	server := Server{}

	server.Type = "FAST HTTP"
	listener, err := reuseport.Listen("tcp6", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	err = fasthttp.Serve(listener, server.HandleFastHTTP)
	//err := fasthttp.ListenAndServe(":9000", server.HandleFastHTTP)
	if err != nil {
		log.Fatal(err)
	}
}
