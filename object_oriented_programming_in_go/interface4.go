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

// Interface Declaration
type Coster interface {
	Cost() float64
}

// Items
type Items []Coster

// START OMIT
// Stringer
func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (t Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", t.Item.String(), t.Cost(), t.discountRate, "%")
}

func (t Items) String() string {
	return fmt.Sprintf("%d items. total: %.0f", len(t), t.Cost())
}

func (o Order) String() string {
	return fmt.Sprintf("Include Tax: %.0f(tax rate: %.2f%s)", o.Cost(), o.taxRate, "%")
}

// END OMIT
// START1 OMIT
func main() {
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	video := Rental{"Interstellar", 1000, 3, Days}
	eventShoes := DiscountItem{
		Item{name: "Women's Walking Shoes", price: 50000, quantity: 3},
		10.00,
	}
	items := Items{shirt, video, eventShoes}
	order := Order{Items: items, taxRate: 10.00}

	fmt.Println(shirt)
	fmt.Println(video)
	fmt.Println(eventShoes)
	fmt.Println(items)
	fmt.Println(order)
}

// END1 OMIT
