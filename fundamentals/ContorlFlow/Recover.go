package main

import (
	"fmt"
	"log"
)

// start line is printed out
// call the panicker function
// inside the func, it calls the panic (panic is called before defer)
// panic function will stop the execution and execute the deferred function
// then it recovers and go back to the main func and prints out "end"
func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// if you want to see the error, do re-panicking
			// panic(err)
		}
	}()
	panic("Somethign happened")
	fmt.Println("Done Panicking")
}
