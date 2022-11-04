package main

import (
	"fmt"
)

func f(c, stop chan int) {
	x := 0
	for {
		fmt.Println("SENT:", x)
		select {
		case c <- x:
			x++
		case v := <-stop:
			fmt.Println("END:", v)
			return
		default:
			fmt.Println("DEFAULT")
		}
	}
}

func main() {
	c := make(chan int)
	stop := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("RECEIVED", <-c)
		}
		stop <- 0
	}() //無名関数の返り値がないということ
	f(c, stop)
}
