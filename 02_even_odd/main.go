package main

import "fmt"

func main() {
	numbers := getNumbers(10)
	even, odd := evenOrOdd(numbers)
	e := countNumbers(even)
	o := countNumbers(odd)

	fmt.Printf("There is %v even numbers:\n\n", e)
	fmt.Print(even)
	fmt.Printf("\n\n\nThere is %v even numbers:\n\n", o)
	fmt.Print(odd)
}
