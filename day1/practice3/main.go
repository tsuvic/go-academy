package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input X:")
	scanner.Scan()

	fmt.Print("Input Y:")
	scanner.Scan()

	x, err := strconv.Atoi(scanner.Text())
	y, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x + y)
}
