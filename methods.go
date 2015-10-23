package main

import (
	"fmt"
)

type rect struct {
	height, width int
}

// This area method has a receiver type of *rect.
func (r *rect) area() int {
	r.width = 20
	r.height = 20
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types.
// Here’s an example of a value receiver.
func (r rect) priem() int {
	r.height = 10
	r.width = 10
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{height: 5, width: 10}

	fmt.Println(r.area())
	fmt.Println(r.priem())

	// Go automatically handles conversion between values and pointers for method calls.
	// You may want to use a pointer receiver type to avoid copying on method calls
	// or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.priem())

	fmt.Println(r.area())
	fmt.Println(r.priem())
}
