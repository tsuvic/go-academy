package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/first", func(w http.ResponseWriter, r *http.Request) {
		body := []byte(`<html><head><title>Go Web Service</title></head><body>トップページ</body><html>`)
		w.Write(body)
	})

	http.HandleFunc("/second/sub", func(w http.ResponseWriter, r *http.Request) {
		body := []byte(`<html><head><title>Go Web Service</title></head><body>サブページ</body><html>`)
		w.Write(body)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗:", err)
	}
}
