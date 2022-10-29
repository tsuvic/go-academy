package main

import "fmt"

type num struct {
	x, y int
}

var (
	a = num{x: 500, y: 100}
)

func (n *num) add(x, y int) int {
	return x + y
}

func (n *num) sub(x, y int) (result int) { //こんなreturn文の書き方もできる
	result = x - y
	return
}

func (n *num) multiply(x, y int) int {
	return x * y
}

func (n *num) divide(x, y int) (result int) {
	result = x / y
	return
}

func main() {
	var n num

	n.x = 200
	n.y = 100
	fmt.Println(n.add(n.x, n.y))
	fmt.Println(n.sub(n.x, n.y))
	fmt.Println(n.multiply(n.x, n.y))
	fmt.Println(n.divide(n.x, n.y))

	fmt.Println(n.add(a.x, a.y))
	fmt.Println(n.sub(a.x, a.y))
	fmt.Println(n.multiply(a.x, a.y))
	fmt.Println(n.divide(a.x, a.y))
}
