package task

import (
	"errors"
)

// Iterator implements the iterator pattern for a slice of tasks
type Iterator struct {
	tasks    []string
	position int
}

// ErrEOF signals a graceful end of iteration
var ErrEOF = errors.New("EOF")

// Next will return the next task in the slice, or ErrEOF once iteration is complete
func (t *Iterator) Next() (int, string, error) {
	t.position++
	if t.position > len(t.tasks) {
		return t.position, "", ErrEOF
	}
	return t.position, t.tasks[t.position-1], nil
}
