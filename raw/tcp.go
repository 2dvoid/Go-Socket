// Implementation of TCP Sockets using raw System Calls

package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	// 1. Create the Socket
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("Socket creation failed: %v", err)
	}
	// Defer closure immediately so it runs no matter where the function exits
	defer unix.Close(fd)

	// 2. Build the address structure
	addr := &unix.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{0, 0, 0, 0},
	}

	// 3. Bind the port to the process 
	err = unix.Bind(fd, addr)
	if err != nil {
		log.Fatalf("Bind failed: %v", err)
	}

	// 4. Start listening 
	err = unix.Listen(fd, 3)
	if err != nil {
		log.Fatalf("Listen failed: %v", err)
	}
	fmt.Println("Server listening on port 8080...")

	// 5. Accept loop
	for {
		nfd, sa, err := unix.Accept(fd)
		if err != nil {
			log.Printf("Accept failed: %v", err)
			continue
		}

		if clientAddr, ok := sa.(*unix.SockaddrInet4); ok {
			ip := clientAddr.Addr
			fmt.Printf("Client connected from %d.%d.%d.%d:%d\n",
				ip[0], ip[1], ip[2], ip[3], clientAddr.Port)
		}

		// 6. Send data
		msg := []byte("Hey there, Stranger!\n")
		_, err = unix.Write(nfd, msg)
		if err != nil {
			log.Printf("Write failed: %v", err)
		}

		// 7. Close the client connection
		// Defer is not used here because we are inside a loop.
        // Defer waits until the surrounding function returns, which would cause a file descriptor leak.
		unix.Close(nfd)
	}
}

