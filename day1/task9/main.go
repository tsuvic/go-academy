package main

import "fmt"

func main() {
	a := map[string]int{
		"Apple": 100,
	}
	a["Banana"] = 50
	a["Orange"] = 200
	a["Strawberry"] = 500
	a["Pineapple"] = 1000

	fmt.Println(a["Apple"])

	fmt.Println("-------------------")

	var ans1 int
	for _, value := range a {
		ans1 += value * 2
	}
	fmt.Println(ans1)

	ans2 := a["Apple"] + a["Strawberry"]
	fmt.Println(ans2)

}
