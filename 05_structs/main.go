package main

import "fmt"

type address struct {
	street string
	city   string
}

type person struct {
	firstName string
	lastName  string
	age       int
	height    int
	children  []person
	address   address
}

func main() {

	// Decalring struct

	// 1.
	andrzej := person{"Andrzej", "Nowak", 29, 186, []person{}, address{}}
	fmt.Println(andrzej)

	// 2.
	marek := person{
		lastName:  "Nowak",
		firstName: "Marek",
		height:    170,
		age:       16,
		children:  nil,
		address: address{
			street: "Szlak",
			city:   "Krakow",
		},
	}
	fmt.Println(marek)

	// 3.
	var jasiek person
	jasiek.firstName = "Jasiek"
	jasiek.lastName = "Meder"
	jasiek.children = []person{andrzej, marek}
	fmt.Println(jasiek)
	fmt.Printf("%+v \n", jasiek)

}
