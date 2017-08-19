package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

const (
	SOCK = "/var/run/appgo.sock"
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

	server.Type = "FAST UNIX"
	err := fasthttp.ListenAndServeUNIX(SOCK, 0777, server.HandleFastHTTP)

	if err != nil {
		log.Fatal(err)
	}
}
