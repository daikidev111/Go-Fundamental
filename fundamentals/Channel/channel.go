package main

import (
	"fmt"
	"math/rand"
	"time"
)


func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(500)
}
func main() {


	// unbuffered channel
	dataChan := make(chan int) // make a channel type int
	// make(chan string, float64...)

	// dataChan <- 789 // both these line have to be executed simultaneously
	// n := dataChan

	// // data channel cannot store data temporarily
	// // it acts as a data transport
	// // Hence, it causes a deadlock
	// fmt.Printf("n = %d", n)

	// Solution to the above
	// you need to have entrace and exit - two portals
	// if there is no destrination -> cause a deadlock
	go func() {
		dataChan <- 789
	}()
	n := <-dataChan
	fmt.Printf("n = %d\n", n)


	// buffered channel
	// you can store a limited nummber of values temporarily

	dataChanBuf := make(chan int, 1)

	dataChanBuf <- 789

	// dataChanBuf <- 123 // what happens if a new int data is added? -> deadlock it is not meant to be added

	a := <- dataChanBuf
	fmt.Printf("%d \n", a)

	// a = <- dataChanBuf
	// fmt.Printf("%d \n", a) // this works and display 123

	// EXAMPLE 1

	// go func() {
	// 	for i:= 0; i < 100; i++ {
	// 		dataChan <- i
	// 	}
	// 	close(dataChan)
	// }() // define the send

	// for b := range dataChan { // define the receive
	// 	fmt.Println(b)
	// }


	// EXAMPLE 2

	go func() {
		for i:=0; i<100; i++ {
			result := DoWork()
			dataChan <- result
		}
		close(dataChan)
	}()

	for c := range dataChan {
		fmt.Println(c)
	}
	
}