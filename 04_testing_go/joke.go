package main

import (
	"fmt"
	"strings"
)

type joke string

func (j joke) tell() {
	fmt.Println("Joke:", j)
}

func (j joke) isPalindrom() bool {
	mid := len(j) / 2
	last := len(j) - 1
	for i := 0; i < mid; i++ {
		if j[i] != j[last-i] {
			return false
		}
	}
	return true
}

func (j joke) reverse() string {
	splitted := strings.Split(string(j), "")
	for i := len(splitted)/2 - 1; i >= 0; i-- {
		opp := len(splitted) - 1 - i
		splitted[i], splitted[opp] = splitted[opp], splitted[i]
	}
	return string(strings.Join(splitted, ""))
}
