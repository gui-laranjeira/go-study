package main

func getNumbers(len int) []int {
	var numbers []int
	for i := 1; i <= len; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func evenOrOdd(numbers []int) ([]int, []int) {
	var even []int
	var odd []int
	for _, number := range numbers {
		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}
	}

	return even, odd
}

func countNumbers(numbers []int) int {
	return len(numbers)
}
