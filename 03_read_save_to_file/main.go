package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	message := "This is a message"

	ioutil.WriteFile("test.txt", []byte(message), 0666)

	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(string(bs))
}
