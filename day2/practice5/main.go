package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input words")
	sc := bufio.NewScanner(os.Stdin)
	data := ""

	for i := 1; sc.Scan(); i++ {
		word := sc.Text()
		fmt.Println(word)
		if word == "quit" {
			break
		}
		data += word + "\n"
	}
	os.WriteFile("day2/practice5/practice.txt", []byte(data), os.ModePerm)
}
