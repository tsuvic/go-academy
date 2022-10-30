package main

import (
	"fmt"
	"time"
)

func f(a string) {
	for i := 0; i < 10; i++ {
		fmt.Println(a, ":", i)
	}
}

func main() {
	go f("a")
	go f("b")
	go f("c")

	time.Sleep(time.Second)
	fmt.Println("done")
}
