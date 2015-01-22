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

type Items []Coster

type Coster interface {
	Cost() float64
}

type Order struct {
	Items   // HL
	taxRate float64
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func (ts Items) Cost() float64 {
	var c float64
	for _, t := range ts {
		c += t.Cost()
	}
	return c
}

func (o Order) Cost() float64 { // HL
	return o.Items.Cost() * (1.0 + o.taxRate/100)
}

// START OMIT
func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	shoes := Item{name: "Sports Shoes", price: 30000, quantity: 2}
	eventShoes := DiscountItem{
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3},
		10.00,
	}
	items := Items{shirt, shoes, eventShoes}
	order := Order{Items: items, taxRate: 10.00}

	fmt.Println("cost of shirt: ", shirt.Cost())
	fmt.Println("cost of shoes: ", shoes.Cost())
	fmt.Println("cost of eventShoes: ", eventShoes.Cost())
	fmt.Println("total cost: ", items.Cost())
	fmt.Println("total cost(included Tax): ", order.Cost())
}

// END OMIT
