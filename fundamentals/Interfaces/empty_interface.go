package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// goof thing for empty interface is that everything  can be casted into an object that has no method on it
	// useful when each other is not type compatible and you need to apply some logic later to figure out exactly what
	// you receieve
	// problem: you cannot do anything with the empty interface becaues there is no method in it
	// usecase: type conversion, use relfect package in order to figure out what kind of an object you are dealing with
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello world from Japan, this is a test"))
		wc.Close()
	}

	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

// Writer interface for writing into bytes
type Writer interface {
	Write([]byte) (int, error)
}

// closer interface for closing the operation
type Closer interface {
	Close() error
}

// composed of writer and closer interfaces
type WriterCloser interface {
	Writer // embedding interface into an interface like struct abstraction
	Closer
}

//struct for creating a concrete methods from the WriterCloser interfaces
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// Concrete method for write
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) // Write out whatever gets sent into the buffered writer closer
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 { // as long as the buffer has more than 8 characters
		_, err := bwc.buffer.Read(v) // read the 8bytes of the buffer
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v)) // print out the string of the 8byte characters
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 { // flush the rest of the buffer that has less than 8 bytes
		data := bwc.buffer.Next(8)          // look for the buffer up to 8 bytes of chars
		_, err := fmt.Println(string(data)) // print out the chracters in the console.
		if err != nil {
			return err
		}
	}
	return nil
}

// construct method to make sure that BufferedWriterCloser is initialized properly
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	} // returning a pointer for data consistency
}
