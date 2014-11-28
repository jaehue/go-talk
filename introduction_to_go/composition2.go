package main

import "fmt"

// Type Declaration (Struct)
type Address struct {
	Number, Street, City, State string
}

// Embedding Types
type Person struct {
	Name string
	Address
}

// START OMIT
// Declaring a Method
func (a *Address) String() string {
	return a.Number + " " + a.Street + "\n" + a.City + ", " + a.State + " " + "\n"
}

func (p *Person) String() string {
	return p.Name + "\n" + p.Address.String() // HL
}

func main() {
	p := Person{
		Name:    "Steve",
		Address: Address{Number: "13", Street: "Main", City: "Gotham", State: "NY"},
	}
	
	fmt.Println(p.String())
	fmt.Println(p.Address.String())
}

// END OMIT
