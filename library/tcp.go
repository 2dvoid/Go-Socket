// Implementation of TCP using "net" library

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept failed: %v", err)
			continue
		}
		
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close() 

	conn.Write([]byte("Hello from TCP!\n"))
	fmt.Printf("%s connected!\n", conn.RemoteAddr().String())
}

