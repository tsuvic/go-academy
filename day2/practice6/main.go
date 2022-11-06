package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.city.shibuya.tokyo.jp/assets/mng/131130_park.csv", nil)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)
	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println(res.Header)
	// fmt.Println(res.Body)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

}
