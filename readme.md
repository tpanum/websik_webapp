# Hints

Run linter
```
golangci-lint run --enable-all main.go
```

Test
```
go test -v
go test -v -race
```

## Example program

``` go
package main

import (
	"fmt"
	"net/http"
)

type server struct {
	i int
}

func (s *server) handleIncrement(w http.ResponseWriter, r *http.Request) {
	s.i += 1
	fmt.Fprintf(w, "ok\n")
}

func (s *server) handleRead(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d\n", s.i)
}

func (s *server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/inc", s.handleIncrement)
	m.HandleFunc("/", s.handleRead)
	return m
}

func main() {
	http.ListenAndServe(":80", nil)
}
```

