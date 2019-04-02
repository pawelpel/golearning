package main

import "fmt"

func main() {
	var ford car = "Ford"
	ford.drive(1)
	ford.drive(2)
	ford.stop()

	polonez := car("Polonez")
	polonez.drive(1)
	polonez.drive(2)
	polonez.stop()

	a1 := createNewHighway1()
	a1 = a1.addCar(ford)
	a1 = a1.addCar(polonez)

	a2 := highway{ford, polonez}
	a2 = a2.addCar(ford)

	fmt.Println("A1:")
	a1.showCars()
	fmt.Println("A2:")
	a2.showCars()
}
