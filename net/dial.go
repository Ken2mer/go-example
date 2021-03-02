package net

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"net/url"
)

// HTTPClient inspired by https://ascii.jp/elem/000/001/276/1276572/
func HTTPClient(rawurl string) (*http.Response, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest("GET", u.String(), nil)
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
