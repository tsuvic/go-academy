package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
			fmt.Println(sum)
		}
	}
	fmt.Println(sum)
}
