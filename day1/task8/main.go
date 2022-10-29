package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [10]int
	fmt.Println(len(a))
	fmt.Printf("%T\n", a)

	b := a[0:10]
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	fmt.Println(len(b))
	fmt.Printf("%T\n", len(b))

	fmt.Printf("----------------- \n")

	var dd [10]int
	d := dd[0:10]
	for i := range d {
		d[i] = i + 1
	}

	fmt.Println(d)
	fmt.Printf("----------------- \n")

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		if c := rand.Intn(9) + 1; c%2 == 0 {
			d = append(d, c)
		} else {
			if len(d) > c { //これがないとクラッシュする
				d = append(d[:c], d[c+1:]...) //...はsliceの要素の値を全て出力してくれるという意味
			}
		}
	}
	fmt.Printf("----------------- \n")
	fmt.Println(d)
}
