package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	SOCK = "/var/run/appgo.sock"
)

type Server struct {
	Type string
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := "Hello World " + s.Type + "\n"
	fmt.Fprint(w, body)
}

func main() {
	server := Server{}

	server.Type = "UNIX SOCKET"
	unix, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.Close()
	err = http.Serve(unix, server)
	if err != nil {
		log.Fatal(err)
	}
}
