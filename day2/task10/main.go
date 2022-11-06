/*
go doc test
*/
package main

//dddd
import (
	"fmt"
)

// bbb
type obj struct {
	x, y int
}

// var (
// 	a = num{x: 500, y: 100}
// )

// aaaa
func newObj() (*obj, error) {
	return &obj{}, nil
}

func (n *obj) add(x, y int) int {
	return x + y
}

func (n *obj) sub(x, y int) (result int) { //こんなreturn文の書き方もできる
	result = x - y
	return
}

func (n *obj) multiply(x, y int) int {
	return x * y
}

func (n *obj) divide(x, y int) (result int) {
	result = x / y
	return
}

func main() {
	var n obj
	n.x = 10
	n.y = 20
	fmt.Println(n.add(n.x, n.y))
	fmt.Println(n.sub(n.x, n.y))
	fmt.Println(n.multiply(n.x, n.y))
	fmt.Println(n.divide(n.x, n.y))

	// fmt.Println(n.add(a.x, a.y))
	// fmt.Println(n.sub(a.x, a.y))
	// fmt.Println(n.multiply(a.x, a.y))
	// fmt.Println(n.divide(a.x, a.y))
}
