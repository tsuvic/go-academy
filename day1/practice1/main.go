package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 10
	y := 20

	z := x + y
	fmt.Println(z)

	var a int32 = 100
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b uint32 = uint32(a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c int64 = int64(a)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	var d float32 = float32(a)
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	i, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	} else {
		fmt.Println("COMLETE CONVERTED:", i)
		fmt.Println("COMLETE CONVERTED:", i, "ERROR VARIABLES:", err)
		fmt.Println("ERROR VARIABLES:", err)
		fmt.Printf("ERROR VARIABLES: %v\n", err)
	}

	j, _ := strconv.Atoi("12345")
	println(j)

	k := strconv.Itoa(12)
	println(k)

}
