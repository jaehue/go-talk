package main

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

type Items []Item

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (ts Items) Cost() float64 {
	var c float64
	for _, t := range ts {
		c += t.Cost()
	}
	return c
}

// START OMIT
// Type Declaration (embedded field)
type Order struct {
	Items   // HL
	taxRate float64
}

// Overriding Methods
func (o Order) Cost() float64 { // HL
	return o.Items.Cost() * (1.0 + o.taxRate/100)
}

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	items := Items{shirt, shoes}
	order := Order{Items: items, taxRate: 10.00}

	fmt.Println("cost of shirt: ", shirt.Cost())
	fmt.Println("cost of shoes: ", shoes.Cost())
	fmt.Println("total cost: ", items.Cost())
	fmt.Println("total cost(included Tax): ", order.Cost())
}

// END OMIT
