// Creational design patterns: Factory pattern example
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get user preference for writer
	kind := "mywriter"
	if len(os.Args) > 1 && os.Args[1] == "stderr" {
		kind = "stderr"
	}
	// Create writer and write some output
	writer, err := NewWriter(kind)
	if err != nil {
		fmt.Print(err)
		return
	}
	writer.Write([]byte("Hello world\n"))
}

// NewWriter is a factory method which instantiates as io.Writer based on the specified kind (mywriter, stderr)
func NewWriter(kind string) (io.Writer, error) {
	switch kind {
	case "mywriter":
		return MyWriter{}, nil
	case "stderr":
		return os.Stderr, nil
	default:
		return nil, fmt.Errorf("NewWriter kind invalid: %s. Must be one of: mywriter, stderr", kind)
	}
}

// MyWriter is an example io.Writer which just prints to stdout
type MyWriter struct{}

// Write implements the only method in the io.Writer interface
func (w MyWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", p)
	return len(p), nil
}
