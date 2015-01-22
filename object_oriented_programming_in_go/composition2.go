package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

// START OMIT
type DiscountItem struct {
	Item
	discountRate float64
}

// Method Declaration
func (t Item) Cost() float64 { // HL
	return t.price * float64(t.quantity) // HL
} // HL

// In Action
func main() {
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	eventShoes := DiscountItem{
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3},
		10.00,
	}

	fmt.Println("cost of shoes: ", shoes.Cost()) // HL
	fmt.Println("cost of eventShoes: ", eventShoes.Cost()) // HL
}

// END OMIT
