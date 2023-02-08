package chan2

import (
	"fmt"
	"time"
)

func Main() {
	fmt.Println("chan2")
	// create channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)

	// receive channel data
	fmt.Println(<-ch)

	msg := <-ch
	fmt.Println(msg)
}
func sendData(ch chan string) {

	// data sent to the channel
	ch <- "Received. Send Operation Successful 1"
	time.Sleep(3 * time.Second)
	ch <- "Received. Send Operation Successful 2"
	time.Sleep(1 * time.Second)
	fmt.Println("No receiver! Send Operation Blocked")
}
