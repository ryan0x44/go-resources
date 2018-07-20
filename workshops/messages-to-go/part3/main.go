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

	// accept a connection and print data received until connection is closed
	connection, err := server.Accept()
	if err != nil {
		log.Printf("Problem with accepting connection: %s", err)
	}
	s := bufio.NewScanner(connection)
	for s.Scan() {
		fmt.Printf("Received: %s\n", s.Text())
	}
	connection.Close()
}
