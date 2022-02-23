package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Printf("%v is big\n", x)
	}

	if x > 100 {
		fmt.Printf("%v is very big\n", x)
	} else {
		fmt.Printf("%v is not that big\n", x)
	}

	if x > 5 && x < 15 {
		fmt.Printf("%v is just right\n", x)
	}

	if x < 20 || x > 30 {
		fmt.Printf("%v is out of range\n", x)
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Printf("%v is more than half of %v\n", a, b)
	}

	// ==================================================

	SwitchStatement()

}

func SwitchStatement() {
	x := 2

	switch x {
	case 1:
		fmt.Printf("One\n")
	case 2:
		fmt.Printf("Two\n")
	case 3:
		fmt.Printf("Three\n")
	default:
		fmt.Printf("Many\n")
	}
}
