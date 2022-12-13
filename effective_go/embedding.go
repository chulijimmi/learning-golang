package main

import (
	"fmt"
	"log"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

// The embedded elements are pointers to structs and of course must be initialized to point to valid
// structs before they can be used. The ReadWriter struct could be written as
type ReadWriterSecond struct {
	reader *Reader
	writer *Writer
}

// Embedding can also be a simple convenience
type Job struct {
	Command string
	*log.Logger
}

func newJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func main() {
	// Composite literal
	command := "ping"
	job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
	fmt.Printf("job: %v\n", job)
}
