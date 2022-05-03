package main

import "fmt"

func main() {

	// Struct
	type point struct {
		x, y int
	}

	fmt.Printf("Binary %b\n", 33)
	fmt.Printf("Unicode %c\n", 33)
	fmt.Printf("Hexadecimal %x\n", 33)

	p := point{1, 2}
	fmt.Printf("%v\n", p)

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	newPerson := Person{"Jeremy", "Morgan", 42}

	fmt.Printf("%T\n", newPerson)

	var isCool = true
	fmt.Printf("Value is %t\n", isCool)

}
