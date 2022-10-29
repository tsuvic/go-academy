package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := func(num int) int {
		return num / 2
	}

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Input X:")
	sc.Scan()
	b, err := strconv.Atoi(sc.Text())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a(b))
	}
}
