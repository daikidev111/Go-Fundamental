package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
const version = "1.0.0"

var pl = fmt.Println // creating an alias for fmt stmt

func main() {

    pl("Enter your name: ")
    reader := bufio.NewReader(os.Stdin)
    name, err := reader.ReadString('\n')
    if err == nil {
        pl("Hello", name)
    } else {
        log.Fatal(err)
    }
}
