package main

import (
	"fmt"
)

func sum(n int, c chan int) {
	total := 0
	for i := 0; i < n; i++ {
		total += i

	}
	c <- total
}

func main() {
	c := make(chan int)
	go sum(10, c)

	result := <-c
	fmt.Println(result)
}
