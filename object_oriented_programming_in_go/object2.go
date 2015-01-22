package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// START OMIT
// Type Declaration
type Items []Item // HL

// Method Declaration
func (ts Items) Cost() float64 { // HL
	var c float64
	for _, t := range ts {
		c += t.Cost()
	}
	return c
}

// In Action
func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	items := Items{shirt, shoes}

	fmt.Println("cost of shirt: ", shirt.Cost())
	fmt.Println("cost of shoes: ", shoes.Cost())
	fmt.Println("total cost: ", items.Cost())
}

// END OMIT
