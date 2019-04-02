package main

import (
	"fmt"
)

func main() {
	newJoke := joke("Baba came to the doctor.")
	newJoke.tell()

	fmt.Println("Palindrom:", newJoke.isPalindrom())
	fmt.Println("Reversed:", newJoke.reverse())

	palindromJoke := joke("lol")
	palindromJoke.tell()

	fmt.Println("Palindrom:", palindromJoke.isPalindrom())
	fmt.Println("Reversed:", palindromJoke.reverse())
}
