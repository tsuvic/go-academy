package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	write()
	read()
}

func write() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Input something you want to output to the file:")
	sc.Scan()
	text := sc.Text()
	err := os.WriteFile("output.txt", []byte(text), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// os.Readfileはファイル全体。bufioは1行ずつ。
func read() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Input the filename you want to read:")

	for i := 0; sc.Scan(); i++ {
		bytes, err := os.ReadFile(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bytes))
	}
}
