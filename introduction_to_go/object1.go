package main

import "fmt"

// START OMIT
// Type Declaration (Struct)
type Rect struct {
	width  int
	height int
}

// Declaring a Method
func (r *Rect) Area() int { // HL
	return r.width * r.height
}

// In Action
func main() {
	r := Rect{width: 10, height: 5}
	fmt.Println("area: ", r.Area()) // HL
}

// END OMIT
