package main

import "fmt"

func main() {
	capitalLetters1 := make(map[string]string)
	fmt.Println(capitalLetters1)

	capitalLetters := map[string]string{
		"r": "R",
		"b": "B",
		"c": "C",
	}
	fmt.Println(capitalLetters)

	capitalLetters["g"] = "G"
	fmt.Println(capitalLetters)

	delete(capitalLetters, "r")
	fmt.Println(capitalLetters)

	printMap(capitalLetters)
}

func printMap(c map[string]string) {
	for lower, capital := range c {
		fmt.Println(lower, " -> ", capital)
	}
}
