package task

// GetExamples will return a slice of example tasks in an iterator
func GetExamples() Iterator {
	return Iterator{
		tasks: []string{
			"say hello",
			"say goodbye",
		},
	}
}
