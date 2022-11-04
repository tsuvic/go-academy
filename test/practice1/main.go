package main

func Sum(n []int) int {
	var total int
	for _, val := range n {
		// fmt.Println(idx)
		total += val
	}
	// fmt.Println(total)
	return total
}

func main() {
	// n = []int {1,2,3,4,5,6,7,8,9,10}
	// n := make([]int, 10)
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Sum(n)
}
