package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, syscall.SIGTERM)

	server := Server{}

	go func() {
		server.Type = "UNIX SOCKET"
		unix, err := net.Listen("unix", SOCK)
		if err != nil {
			log.Fatal(err)
		}
		http.Serve(unix, server)
	}()

	go func() {
		server.Type = "HTTP"
		http.Handle("/", server)
		if err := http.ListenAndServe(":9000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	<-sigchan

	if err := os.Remove(SOCK); err != nil {
		log.Fatal(err)
	}
}
