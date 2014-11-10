package main

import "fmt"

// Type Declaration
type Rect struct {
	width  int
	height int
}
type Rects []Rect

// Declaring a Method
func (r Rect) Area() int {
	return r.width * r.height
}

func (rs Rects) Area() int {
	var a int
	for _, r := range rs {
		a += r.Area()
	}
	return a
}

// START OMIT
// Interface Declaration (Struct)
type Shaper interface {
	Area() int // HL
}

// Using Interface as Param Type
func Describe(s Shaper) {
	fmt.Println("Area is:", s.Area())
}

// In Action
func main() {
	r := Rect{width: 10, height: 5}
	x := Rect{width: 7, height: 10}
	rs := Rects{r, x}

	Describe(r)  // HL
	Describe(x)  // HL
	Describe(rs) // HL
}

// END OMIT
