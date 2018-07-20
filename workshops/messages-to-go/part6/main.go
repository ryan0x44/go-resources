package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// parse command-line flags
	dial := flag.String("dial", "", "dial address")
	port := flag.String("port", "", "port to serve on")
	flag.Parse()
	if *dial == "" || *port == "" {
		flag.Usage()
		return
	}
	fmt.Printf("Dial address: %s\n", *dial)

	// handle input concurrently and dial server/send message for every line of input
	in := bufio.NewScanner(os.Stdin)
	go func() {
		for in.Scan() {
			fmt.Printf("You said: %s\n", in.Text())
			c, err := net.Dial("tcp", *dial)
			if err != nil {
				log.Printf("Unable to connect to %s - %s\n", *dial, err)
				continue
			}
			_, err = c.Write(in.Bytes())
			if err != nil {
				log.Printf("Unable to write to connection\n")
				continue
			}
			log.Printf("Sent to %s\n", *dial)
			c.Close()
		}
	}()

	// listen for connections
	server, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Unable to start server")
	}
	defer server.Close()
	log.Printf("Listening on localhost:" + *port)

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
