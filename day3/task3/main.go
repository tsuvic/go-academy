package main

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", handleIndex)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("day3/task3/index.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	template := Template{"タイトル", "本文"}

	if err := t.Execute(w, template); err != nil {
		log.Printf("テンプレート埋め込みに失敗 %v", err)
	}

}
