package main

import "fmt"

// START1 OMIT
type Rect struct {
	width  int
	height int
}
type Rects []Rect

func (r Rect) Area() int {	// HL
	return r.width * r.height
}

func (rs Rects) Area() int {	// HL
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
// END1 OMIT

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
