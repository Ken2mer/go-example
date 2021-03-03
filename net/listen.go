package net

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
)

// HTTPServer inspired by https://ascii.jp/elem/000/001/276/1276572/
func HTTPServer(port string) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ln, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return err
	}

	go func() {
		s := <-c
		fmt.Println("Got signal:", s)
		ln.Close()
	}()

	fmt.Println("Server is running at localhost:" + port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// リクエストを読み込む
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				fmt.Println(err)
				return
			}
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(dump))

			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)

			conn.Close()
		}()
	}
}
