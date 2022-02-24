package main

import "fmt"

// OOP

type Trade struct {
	// if property starts with captal letter that's poblic otherwise is private
	Symbol string  // stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade Price
	Buy    bool    // true if buy trade. false if sell trade
}

func main() {
	t1 := Trade{"APPLE", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Printf("Price: %+v\n", t1.Price)

	fmt.Println("=======================================")

	t2 := Trade{
		Volume: 15,
		Symbol: "MICROSOFT",
		Price:  99.97,
		Buy:    false, // trailing come is nandatory
	}

	fmt.Println(t2)
	fmt.Printf("%+v\n", t2)

	fmt.Println("=======================================")

	t3 := Trade{}
	fmt.Println(t3)
	fmt.Printf("%+v\n", t3)
}
