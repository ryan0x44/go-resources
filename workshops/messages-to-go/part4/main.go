package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// listen for connections
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Unable to start server")
	}
	defer server.Close()
	log.Printf("Listening on localhost:8080")

	// accept connections and print received messages
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Printf("Problem with accepting connection: %s", err)
		}
		go func() {
			s := bufio.NewScanner(connection)
			for s.Scan() {
				fmt.Printf("Received: %s\n", s.Text())
			}
			connection.Close()
		}()
	}
}
