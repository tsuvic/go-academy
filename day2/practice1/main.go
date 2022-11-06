package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ans := make(map[int][]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 30; i++ {
		num := rand.Intn(100)
		switch {
		case num%13 == 0:
			ans[3] = append(ans[3], num)
		case num%27 == 0:
			ans[3] = append(ans[3], num)
		case num%37 == 0:
			ans[3] = append(ans[3], num)
		case num%2 == 1:
			ans[2] = append(ans[2], num)
		case num%2 == 0:
			ans[1] = append(ans[1], num)
		}
	}
	fmt.Println(ans)
}
