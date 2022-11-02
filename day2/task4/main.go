package main

import (
	"fmt"
	"time"
)

func main() {
	go one()
	go two()
	go three()
	time.Sleep(time.Second * 4)
	fmt.Println("done")
}

func one() {
	now := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Printf("経過時間: %vms\n", time.Since(now).Milliseconds())
}

func two() {
	now := time.Now()
	time.Sleep(time.Second * 2)
	fmt.Printf("経過時間: %vms\n", time.Since(now).Milliseconds())
}

func three() {
	now := time.Now()
	time.Sleep(time.Second * 3)
	fmt.Printf("経過時間: %vms\n", time.Since(now).Milliseconds())
}
