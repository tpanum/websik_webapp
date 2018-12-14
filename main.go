package main

import (
	"fmt"
	"net/http"
	"sync"
)

type server struct {
	m sync.Mutex
	i int
}

func (s *server) handleRead(w http.ResponseWriter, r *http.Request) {
	s.m.Lock()
	fmt.Fprintf(w, "%d", s.i)
	s.m.Unlock()
}

func (s *server) handleInc(w http.ResponseWriter, r *http.Request) {
	s.m.Lock()
	s.i = s.i + 1
	s.m.Unlock()
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
