package main

import "fmt"

type developer interface {
	getLanguage() string
}

type pythonDev struct{}
type golangDev struct{}

func main() {
	pyDev := pythonDev{}
	goDev := golangDev{}

	showLanguage(pyDev)
	showLanguage(goDev)
}

func showLanguage(d developer) {
	fmt.Println(d.getLanguage())
}

func (pythonDev) getLanguage() string {
	return "Python"
}

func (golangDev) getLanguage() string {
	return "Golang"
}
