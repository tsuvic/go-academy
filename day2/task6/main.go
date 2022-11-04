package main

import (
	"fmt"
	"time"
)

func send(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Send", i)
		c <- i
		// time.Sleep(time.Microsecond * 10)
	}
}

func receive(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		n := <-c
		fmt.Println("Received", n)
	}
}

func main() {
	c := make(chan int, 1)
	go send(c)
	go receive(c)

	//closed channelにキュー送付を確認する
	c <- 1

	time.Sleep(time.Second * 1)

	//キュー取得の空振り確認用
	c <- 1
	fmt.Println("AFTER", <-c)
	c <- 2
	fmt.Println("AFTER", <-c)

	//バッファー溢れ確認用
	d := make(chan int, 1)
	d <- 1
	// d <- 2　ここを外す

	//バッファーチャンネルの確認用
	e := make(chan int, 1)
	e <- 1
	v, ok := <-e
	if !ok {
		panic(ok)
	}
	println(v)

	// e <- 1
	// v2, ok2 := <-e
	// v3, ok3 := <-e
	// if !ok2 || !ok3 {
	// 	panic("FAIL")
	// }
	// println(v2)
	// println(v3)

}
