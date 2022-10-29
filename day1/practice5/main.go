package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		number := int64(rand.Intn(1000))
		fmt.Print(number)
		if number == 7 {
			fmt.Printf("You got number %v\n", number)
			fmt.Printf("You got number %v\n", number)
			break
		}
	}
}
