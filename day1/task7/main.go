package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	// fmt.Println(a)
	// b := make([]int, 10)
	// fmt.Println(b)
	// c := make([]int, 5, 10)
	// fmt.Println(c)
	// var d [3]int
	// fmt.Println(d)

	fmt.Println(a)
	fmt.Println(reverse(a))

}

func reverse(numbers []int) []int {
	newNumbers := make([]int, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}
