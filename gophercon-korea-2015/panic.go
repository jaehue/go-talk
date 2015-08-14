package main

import "fmt"

// START OMIT
func main() {
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}

// END OMIT
