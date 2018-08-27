// Creational design patterns: Dependency Injection (Constructor) pattern
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := NewMyService(os.Stderr)
	s.WriteHello("world")
}

// MyService is an example service with an io.Writer dependency
type MyService struct {
	writer io.Writer
}

// NewMyService is a constructor for MyService which accepts an io.Writer dependency
func NewMyService(writer io.Writer) MyService {
	return MyService{
		writer: writer,
	}
}

// WriteHello demonstrates writing a string to the io.Writer dependency
func (s *MyService) WriteHello(m string) {
	fmt.Fprintf(s.writer, "Hello %s\n", m)
}
