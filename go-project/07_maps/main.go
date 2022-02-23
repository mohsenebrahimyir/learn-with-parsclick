package main

import "fmt"

func main() {
	bar := map[string]float64{
		"AMAZON":    1700.0,
		"GOOGLE":    1170.6,
		"MICROSOFT": 100.5,
	}

	// Get a value
	fmt.Println(bar["MICROSOFT"])

	fmt.Println("==============================")

	// Get zero value if not found
	fmt.Println(bar["APPLE"])

	fmt.Println("==============================")

	// Use two value to see if found
	value, ok := bar["APPLE"]

	if ok {
		fmt.Println("APPLE not found")
	} else {
		fmt.Println(value)
	}

	// Set
	bar["APPLE"] = 350.0
	fmt.Println(bar)

	fmt.Println("==============================")

	// Delete
	delete(bar, "AMAZON")
	fmt.Println(bar)

	fmt.Println("==============================")

	// Single value "for" is on keys
	for key := range bar {
		fmt.Println(key)
	}

	fmt.Println("==============================")

	// Double value "for" is key, value
	for key, value := range bar {
		fmt.Printf("%s - %.2f\n", key, value)
	}

	fmt.Println("==============================")

	// Double value "for" is on value and ignore key
	for _, value := range bar {
		fmt.Println(value)
	}
}
