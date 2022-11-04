package main

import "fmt"

func sum(n int, c chan int) {
	total := 0
	for i := 0; i < n; i++ {
		total += i
	}
	c <- total
}

func main() {
	c := make(chan int)
	go sum(0, c)
	go sum(1, c)
	go sum(1000000000, c)

	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println(result1, result2, result3)
	fmt.Println(&result1, &result2, &result3)
	//FIFO
}
