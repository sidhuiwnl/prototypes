package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 4096)
	n, _ := conn.Read(buffer)

	fmt.Println("---- Incoming Request ----")
	request := string(buffer[:n])
fmt.Println(strings.Split(request, "\n")[0])

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 40\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"Hello, World!"

	conn.Write([]byte(response))
}