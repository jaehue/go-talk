package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

type DiscountItem struct {
	Item
	discountRate float64
}

// START OMIT

// Method Declaration
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t DiscountItem) Cost() float64 { // HL
	return t.Item.Cost() * (1.0 - t.discountRate/100) // HL
} // HL

// In Action
func main() {
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	eventShoes := DiscountItem{
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3},
		10.00,
	}

	fmt.Println("cost of shoes: ", shoes.Cost())
	fmt.Println("cost of eventShoes: ", eventShoes.Cost())
}

// END OMIT
