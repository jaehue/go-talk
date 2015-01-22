package main

import "fmt"

// Item
type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// Rental
// START1 OMIT
type Rental struct {
	name         string
	feePerDay    float64
	periodLength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

// END1 OMIT

// START2 OMIT
func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

// Interface Declaration
type Coster interface { // HL
	Cost() float64 // HL
} // HL

func DisplayCost(c Coster) { // HL
	fmt.Println("cost: ", c.Cost())
}

// In Action
func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	video := Rental{"Interstellar", 1000, 3, Days}

	fmt.Printf("[%v] ", shirt.name)
	DisplayCost(shirt) // HL
	fmt.Printf("[%v] ", video.name)
	DisplayCost(video) // HL
}

// END1 OMIT
// END2 OMIT
