package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Processing the request")
	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("Waiting for client to connect")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Client connected")
		go do(conn)
	}
}
