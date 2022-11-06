package main

import (
	"fmt"
	"time"
)

func main() {
	ans := "2021年10月31日(日)2時0分 Sunday, October 31, 2021 2:00 AM"
	println(ans)
	fmt.Println("JAPAN", time.Now())

	New_York, _ := time.LoadLocation("America/New_York")
	fmt.Println("NEW YORK", time.Now().In(New_York))

	// HongKong, _ := time.LoadLocation("America/New_York")
	// fmt.Println("HongKong", time.Now().In(HongKong))

	// London, _ := time.LoadLocation("America/New_York")
	// fmt.Println("NEW YORK", time.Now().In(London))

}
