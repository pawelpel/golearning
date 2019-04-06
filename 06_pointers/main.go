package main

import "fmt"

type car struct {
	name  string
	model string
}

// Pass By Value - what means that in function with reciever we are updating
// new copy of given struct.
func (c car) setName(newName string) {
	c.name = newName
}

func (pointerToCar *car) setModel(newModel string) {
	(*pointerToCar).model = newModel
}

func main() {
	ford := car{name: "Trabant", model: "Ford"}
	ford.setName("Voyager")

	fmt.Printf("%+v\n", ford)

	pointerToFord := &ford
	pointerToFord.setModel("Opel")

	fmt.Printf("%+v\n", ford)

	// Go do the shortuct and automatically change `ford` into `pointerToFord`
	// so we can omit line `pointerToFord := &ford`
	ford.setModel("Mercedes")

	fmt.Printf("%+v\n", ford)

	// Gotchas
	mySlice := []string{"Elo", "mordo"}
	updateSlice(mySlice)
	// it's updated cause mySlice is structure with pointer to an array
	// and we pass a value so that structure is copied but it still
	// points to the same array
	fmt.Println(mySlice)

	// VALUE TYPE
	// int, float, string, bool, struct <- use pointers to change values inside a function

	// REFERENCE TYPE
	// slice, map, channel, pointer, functions <- don't worry about pointer with these

}

func updateSlice(s []string) {
	s[0] = "Siema"
}
