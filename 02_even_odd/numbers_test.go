package main

import (
	"testing"
)

func TestGetNumbers(t *testing.T) {
	n := getNumbers(2000)
	if len(n) != 2000 {
		t.Errorf("Expected 2000 numbers but got %v", len(n))
	}

	if n[0] != 1 {
		t.Errorf("First number expected were 1, but got %v", n[0])
	}

	if n[len(n)-1] != 2000 {
		t.Errorf("Last number expected were 2000, but got %v", n[len(n)-1])
	}
}

func TestEvenOrOdd(t *testing.T) {
	n := getNumbers(200)
	e, o := evenOrOdd(n)

	for _, number := range e {
		if number%2 != 0 {
			t.Errorf("Expected even numbers, but got %v", number)
			return
		}
	}
	for _, number := range o {
		if number%2 == 0 {
			t.Errorf("Expected odd numbers, but got %v", number)
		}
	}
}
