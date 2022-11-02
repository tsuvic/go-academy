package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://go.dev/solutions", nil)
	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	// fmt.Println(io.ReadAll(res.Body))
	fmt.Println(string(body))

	fmt.Println(res.StatusCode)
	fmt.Println(res.Header.Get("Content-Type"), "ヘッダー一部")
	fmt.Println(res.Header, "ヘッダー全部")
	fmt.Println(res.ContentLength, "長さ")

	err := os.WriteFile("test.html", []byte(body), os.ModePerm)
	log.Fatal(err)
}
