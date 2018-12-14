package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
)

type Server interface {
	Routes() http.Handler
}

func TestServer(t *testing.T) {
	// instantiate this !!!
	var s Server
	s = &server{}

	tt := []struct {
		name       string
		increments int
		diff       int
	}{
		{name: "+4", increments: 4, diff: 4},
		{name: "+0", increments: 0, diff: 0},
		{name: "+300", increments: 300, diff: 300},
		{name: "+1", increments: 1, diff: 1},
	}

	ts := httptest.NewServer(s.Routes())
	defer ts.Close()

	// read performs a HTTP request to server with path "/" returns count
	read := func() (int, error) {
		result, err := http.Get(ts.URL + "/")
		if err != nil {
			return 0, err
		}

		count, err := ioutil.ReadAll(result.Body)
		result.Body.Close()
		if err != nil {
			return 0, err
		}

		return strconv.Atoi(strings.TrimSpace(string(count)))
	}

	// increment performs a HTTP request to server with path "/inc"
	increment := func(wg *sync.WaitGroup) error {
		defer wg.Done()

		result, err := http.Get(ts.URL + "/inc")
		if err != nil {
			return err
		}

		content, err := ioutil.ReadAll(result.Body)
		result.Body.Close()
		if err != nil {
			return err
		}

		trimmedContent := strings.TrimSpace(string(content))
		if trimmedContent != "ok" {
			return fmt.Errorf("response was not \"ok\"")
		}

		return nil
	}

	for i, _ := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {

			count, err := read()
			if err != nil {
				t.Fatalf("error when reading: %s", err)
			}

			var wg sync.WaitGroup
			wg.Add(tc.increments)
			for i := 0; i < tc.increments; i++ {
				go increment(&wg)
			}
			wg.Wait()

			postCount, err := read()
			if err != nil {
				t.Fatalf("error when reading: %s", err)
			}

			if readDiff := postCount - count; readDiff != tc.diff {
				t.Fatalf("unexpected diff (expected: %d): %d", tc.diff, readDiff)
			}

		})
	}

}
