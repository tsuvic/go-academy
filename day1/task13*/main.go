package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, y := in()
	sum(x, y)
}

func in() (x, y int) {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Input number :")
	sc.Scan()
	num1, _ := strconv.Atoi(sc.Text())

	fmt.Println("Input number :")
	sc.Scan()
	num2, _ := strconv.Atoi(sc.Text())

	return num1, num2
}

func sum(x, y int) {
	fmt.Printf("The answer is %v\n", x+y)
}
