package main

import (
	"fmt"
	"time"
)

func main() {
	// initialize a counter
	i := 0

	for {
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
	}
}
