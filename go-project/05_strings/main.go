package main

import "fmt"

func main() {
	book := "Yare dabestanie man"
	//       0123456789...
	fmt.Println(book)

	// string length
	fmt.Println(len(book))

	fmt.Printf("book[0] = %v, type = %T\n", book[0], book[0])
	// (strings are immutable)
	// book[0] = 110

	// slice (start, end) 0 based, half empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concationate string
	fmt.Println("May " + book)

	// Multiline
	poem1 := `
	Yare dabestany man
	Ba mano hamrah mani
	`
	poem2 := `
Yare dabestany man
Ba mano hamrah mani
	`
	fmt.Println(poem1)
	fmt.Println(poem2)

}
