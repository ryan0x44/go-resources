// Behavioral design patterns: Iterator pattern example
package main

import (
	"log"

	"github.com/ryan0x44/go-resources/design-patterns/iterator/task"
)

func main() {
	tasks := task.GetExamples()
	for {
		i, t, err := tasks.Next()
		if err == task.ErrEOF {
			log.Printf("Done")
			break
		}
		if err != nil {
			log.Fatalf("Unknown error: %s", err)
		}
		log.Printf("Task %d: %s\n", i, t)
	}
}
