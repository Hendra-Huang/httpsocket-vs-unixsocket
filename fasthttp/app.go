package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
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
	err := fasthttp.ListenAndServe(":9000", server.HandleFastHTTP)
	if err != nil {
		log.Fatal(err)
	}
}
