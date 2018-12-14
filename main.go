package main

import (
	"fmt"
	"net/http"
)

type server struct {
	i int
}

func (s *server) handleRead(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", s.i)
}

func (s *server) handleInc(w http.ResponseWriter, r *http.Request) {
	s.i = s.i + 1
	fmt.Fprintf(w, "ok")
}

func (s *server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", s.handleRead)
	m.HandleFunc("/inc", s.handleInc)

	return m
}

func main() {

}
