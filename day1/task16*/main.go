package main

import "fmt"

func main() {
	s := "Hello, Banana"
	fmt.Println(s[(len(s) - 3):])
}
