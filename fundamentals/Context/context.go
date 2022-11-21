package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)


var pl = fmt.Println

/*
	context is a way to add a timeout or cancellation to a goroutine
	if it takes a long time to execute, it is a way to cancel the process
*/

func main() {

	timeoutContext, cancel := context.withTimeout(context.Background(), time.Millisecond()*50)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	imageData, err := iovtil.ReadAll(res.Body)
	pl("Downloaded image of size %d\n", len(imageData))
}

