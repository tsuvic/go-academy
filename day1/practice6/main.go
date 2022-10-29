package main

import "fmt"

func main() {

	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array1)

	var array2 [100]int
	for key := range array2 {
		{
			array2[key] = key + 1
			// fmt.Println(array2[key])
		}
	}
	fmt.Println(array2)

	var array3 []int
	for key := range array3 {
		{
			array3[key] = key + 1
			// fmt.Println(array2[key])
		}
	}
	fmt.Println(array3)

	for i := 0; i < len(array2); i++ {
		switch {
		case array2[i]%15 == 0:
			fmt.Println(i, "hogefuga")
		case array2[i]%5 == 0:
			fmt.Println(i, "hoge")
		case array2[i]%3 == 0:
			fmt.Println(i, "fuga")
		default:
			fmt.Println(i, "no")
		}
	}

	//fallthroughしてしまうと、次のcaseの条件判定は行われず、無条件で処理が実行される
	for i := 0; i < len(array2); i++ {
		var ans string
		switch {
		case array2[i]%5 == 0:
			ans += "hoge"
			fallthrough
		case array2[i]%3 == 0:
			ans += "fuga"
		default:
			ans += "no"
		}
		fmt.Printf("Value: %v  String: %v\n", array2[i], ans)
	}
}
