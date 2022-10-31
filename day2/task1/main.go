package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	write()
}

func write() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Input something you want to output to the file:")
	sc.Scan()
	text := sc.Text()
	err := os.WriteFile("output.txt", []byte(text), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func read() {

}
