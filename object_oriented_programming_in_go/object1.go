package main

import "fmt"

// Type Declaration
type Item struct { // HL
	name     string
	price    float64
	quantity int
}

// Method Declaration
func (t Item) Cost() float64 { // HL
	return t.price * float64(t.quantity)
}

// In Action
func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	fmt.Println("cost: ", shirt.Cost()) // HL
}
