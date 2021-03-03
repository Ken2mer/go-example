package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall"
	"testing"
	"time"
)

func TestHTTPServer(t *testing.T) {
	port := "8080"
	errCh := make(chan error)
	go func() {
		errCh <- HTTPServer(port)
	}()
	time.Sleep(time.Second * 1)

	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if status != 200 {
		t.Fatalf("expected: 200, got: %d", resp.StatusCode)
	}
	if string(body) != "Hello World\n" {
		t.Fatalf("expected: Hello World\n, got: %s", body)
	}

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	fmt.Println(<-errCh)
}
