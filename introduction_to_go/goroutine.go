package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(msg string, delay time.Duration) {
	for {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}

func main() {
	go f("A--", 300*time.Millisecond)
	go f("-B-", 500*time.Millisecond)
	go f("--C", 1100*time.Millisecond)
	time.Sleep(10 * time.Second)
}

// END OMIT
