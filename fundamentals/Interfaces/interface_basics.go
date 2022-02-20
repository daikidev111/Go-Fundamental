package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

// this is how u declare an interface
// it does not describe data
// it stores the method definition
type Writer interface {
	Write([]byte) (int, error)
}

// we are implicitly implementing the interface
// by creating a method on our console writer ithat has the
// signature of a writer interface
type ConsoleWriter struct {
}

// concrete implmementation of the Write method from the Writer interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
