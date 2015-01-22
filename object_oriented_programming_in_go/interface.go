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

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

// DiscountItem
type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

// Order
type Order struct {
	Items
	taxRate float64
}

func (o Order) Cost() float64 { // HL
	return o.Items.Cost() * (1.0 + o.taxRate/100)
}

func (ts Items) Cost() float64 {
	var c float64
	for _, t := range ts {
		c += t.Cost()
	}
	return c
}

func DisplayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

// START OMIT
// Interface Declaration
type Coster interface {
	Cost() float64
}

// Items
type Items []Coster // HL

func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3},
		10.00,
	}
	items := Items{shirt, video, eventShoes}
	order := Order{Items: items, taxRate: 10.00}

	DisplayCost(shirt)
	DisplayCost(video)
	DisplayCost(eventShoes)
	DisplayCost(items)
	DisplayCost(order)
}

// END OMIT
