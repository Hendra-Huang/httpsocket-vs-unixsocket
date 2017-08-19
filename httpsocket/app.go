package main

import (
	"fmt"
	"log"
	"net/http"
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

	server.Type = "HTTP"
	http.Handle("/", server)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
