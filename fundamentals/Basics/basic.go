package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

var pl = fmt.Println

func main() {
	// variable declaration
	// var vName string = "Daiki"
	// var v1, v2 = 1.2, 1.3

	// how to cast
	cV1 := 1.5
	cV2 := int(cV1)
	pl(reflect.TypeOf(cV2))
	cV3 := "50000"
	cV4, err := strconv.Atoi(cV3) // convert from string to int
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000
	cV6 := strconv.Itoa(cV5)
	pl(cV6, reflect.TypeOf(cV6)) // convert from int to string

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil { pl(cV8) } else { pl(err) }

	// data type checking
	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf(3.14))
	// pl(reflect.TypeOf(true))
	// pl(reflect.TypeOf("Hello"))


	// date and time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
	
	// generate a random numnber
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)
}