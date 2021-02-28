package net

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"strings"
)

// HTTPClient inspired by https://ascii.jp/elem/000/001/276/1276572/
func HTTPClient(url string) (*http.Response, error) {
	conn, err := net.Dial("tcp", strings.Split(url, "://")[1])
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Write(conn)

	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		log.Fatal(err)
	}
	return response, err
}
