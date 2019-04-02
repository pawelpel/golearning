package main

import "testing"

func TestIsPalindrom(t *testing.T) {
	joke1 := joke("Baba ma kota, a sierotka Marysia.")

	if joke1.isPalindrom() != false {
		t.Errorf("Expected joke which isn't a palindrom.")
	}

	joke2 := joke("Baba abaB")
	if joke2.isPalindrom() == false {
		t.Errorf("Expected joke which is a palindrom.")
	}

}

func TestRevserse(t *testing.T) {
	joke1 := joke("baba")

	if joke1.reverse() != "abab" {
		t.Errorf("Expected: 'abab', got: %v", joke1)
	}
}

func TestTell(t *testing.T) {
	joke("Hello").tell()
}
