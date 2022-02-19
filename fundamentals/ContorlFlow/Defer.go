package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("============== 1st =============")
	// defer acts like stack. It pops off from the top of the stack
	// which is "End" in this case
	defer fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("End")

	fmt.Println("============== 2nd =============")

	// this throws an error beause we closed the response body before we were done reading it.
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res.Body.Close()
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// Main pattern for usijng defer is to put after the opening of a resource
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

}
