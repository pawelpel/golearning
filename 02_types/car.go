package main

import (
	"fmt"
	"strconv"
)

type car string

func (c car) drive(direction int) {
	fmt.Println(c, "driving in direction:", direction, "...")
}

func (c car) stop() {
	fmt.Println("Stopped!")
}

type highway []car

func (h highway) addCar(newCar car) highway {
	return append(h, newCar)
}

func (h highway) showCars() {
	for number, carObject := range h {
		fmt.Println(strconv.Itoa(number+1)+".", carObject)
	}
}

func createNewHighway1() highway {
	var newHighway highway
	return newHighway
}

func createNewHighway2() highway {
	return highway{}
}
