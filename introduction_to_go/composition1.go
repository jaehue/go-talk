package main

import "fmt"

// START OMIT
// Type Declaration (Struct)
type Address struct {
	Number, Street, City, State string
}

// Embedding Types
type Person struct {
	Name    string
	Address // HL
}

func main() {
	// Declare using Composite Literal
	p := Person{
		Name:    "Steve",
		Address: Address{Number: "13", Street: "Main", City: "Gotham", State: "NY"}, // HL
	}
	
	fmt.Println(p)
	fmt.Println(p.Address)
}

// END OMIT
