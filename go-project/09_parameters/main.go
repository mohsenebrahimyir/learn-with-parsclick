package main

import "fmt"

func DoubleAt(values []int, i int) { // slices, maps pass by references
	values[i] *= 2
}

func Double(n int) { // passed by value (int)
	n *= 2
}

func DoublePtr(n *int) { // pointer
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	DoubleAt(values, 2)
	fmt.Println(values)

	fmt.Println("=============================")

	val := 10
	Double(val)
	fmt.Println(val)

	fmt.Println("=============================")

	DoublePtr(&val)
	fmt.Println(val)
}
