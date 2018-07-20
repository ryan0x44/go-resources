package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "", "host to connect to")
	flag.Parse()
	if *host == "" {
		flag.Usage()
		return
	}
	fmt.Printf("Host: %s\n", *host)
}
