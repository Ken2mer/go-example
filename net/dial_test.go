package net

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPClient(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	resp, err := HTTPClient(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if status != 200 {
		t.Fatalf("expected: 200, got: %d", resp.StatusCode)
	}
	if string(body) != "Hello, client\n" {
		t.Fatalf("expected: Hello, client\n, got: %s", body)
	}
}
