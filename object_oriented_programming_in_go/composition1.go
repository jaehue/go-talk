package main

import "fmt"

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// START OMIT

// Type Declaration
type Item struct {
	name     string
	price    float64
	quantity int
}

type DiscountItem struct {
	Item         // HL
	discountRate float64
}

// In Action
func main() {
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	eventShoes := DiscountItem{ // HL
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3}, // HL
		10.00, // HL
	} // HL

	fmt.Println("shoes: ", shoes)
	fmt.Println("eventShoes: ", eventShoes)
}

// END OMIT
