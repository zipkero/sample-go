package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "hello, world")
			},
		))
	return ts
}

func TestFetchRemoveResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()
	excepted := "hello, world"
	data, err := fetchRemoveResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if excepted != string(data) {
		t.Errorf("expected response to be: %s, goet: %s", excepted, data)
	}
}
