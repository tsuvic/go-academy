package main

import (
	"log"
	"net/http"
)

func main() {
	body := []byte(`<html><head><title>Go Web Service</title></head><body>トップページ</body><html>`)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(body)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗:", err)
	}
}
