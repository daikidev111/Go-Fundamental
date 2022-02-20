package main

import (
	"fmt"
)

// Summary:When implementing an interface, if I use a value type
// for interface declaration, the methods that implement the intergace have to
// have all value receivers.
// If the interface with a pointer, then the methods that implement the interface have to
// have all pointer receivers
func main() {
	var i interface{} = 0
	switch i.(type) { // type switch
	case int:
		fmt.Println("i is an integer")

	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I dont know what i is")
	}

	// var wc WriterCloser = myWriterCloser{}
	// wc is holding the myWriterCloser interface value
	var wc WriterCloser = &myWriterCloser{}

	fmt.Println(wc)
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

type myWriterCloser struct{}

// Cause an error
// ./interface_type_conversion_another.go:19:6:
// cannot use myWriterCloser{} (type myWriterCloser)
// as type WriterCloser in assignment:
// myWriterCloser does not implement WriterCloser
// (Write method has pointer receiver) <-- this is important
// Solution:
// make it a pointer receiever
func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// Solution:
// make it a pointer receiever
func (mwc *myWriterCloser) Close() error {
	return nil
}
