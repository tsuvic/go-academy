package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.PostForm

	id, ok := form["id"]
	if ok {
		searchById(id[0])
	}

	title, ok := form["title"]
	if ok {
		csvSearch(title[0])
	}

	if !ok {
		fmt.Println("no params")
	}

}

func searchById(id string) {
	f, err := os.Open("task14/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	var index int

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	}

}
