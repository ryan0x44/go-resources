// Creational design patterns: Dependency Injection (Setter) pattern
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	s := NewMyService()
	s.SetWriter(os.Stderr)
	s.WriteHello("world")
}

// MyService is an example service with an io.Writer dependency
type MyService struct {
	writer io.Writer
}

// NewMyService is a constructor for MyService
func NewMyService() MyService {
	return MyService{}
}

// SetWriter can be used to specify the io.Writer dependency for MyService
func (s *MyService) SetWriter(w io.Writer) {
	s.writer = w
}

// WriteHello demonstrates writing a string to the io.Writer dependency
func (s *MyService) WriteHello(m string) {
	fmt.Fprintf(s.writer, "Hello %s\n", m)
}
