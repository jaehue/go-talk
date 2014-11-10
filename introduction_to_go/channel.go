package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(msg string, delay time.Duration, ch chan string) {
	for {
		ch <- msg
		time.Sleep(delay)
	}
}

func main() {
	ch := make(chan string)
	go f("A--", 300*time.Millisecond, ch)
	go f("-B-", 500*time.Millisecond, ch)
	go f("--C", 1100*time.Millisecond, ch)

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-ch)
	}
}

// END OMIT
