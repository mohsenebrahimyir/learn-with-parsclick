package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func DivieModular(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	x := 34
	y := 7

	sum := Add(x, y)
	fmt.Println(sum)

	div, mod := DivieModular(x, y)
	fmt.Printf("div = %d, mod = %d\n", div, mod)
}
