package main

import "fmt"

// START1 OMIT
type Rect struct {
	width  int
	height int
}
type Rects []Rect

func (r Rect) Area() int {	// HLinterface
	return r.width * r.height
}

func (rs Rects) Area() int {	// HLinterface
	var a int
	for _, r := range rs {
		a += r.Area()
	}
	return a
}

// START OMIT
// Interface Declaration
type Shaper interface { // HLinterface
	Area() int // HLinterface
} // HLinterface
// END1 OMIT

// Using Interface as Param Type
func Describe(s Shaper) { // HL
	fmt.Println("Area is:", s.Area()) // HL
} // HL

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
