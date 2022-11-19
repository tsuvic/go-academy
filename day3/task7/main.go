package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StatusJSON struct {
	Date      string      //`json:"date"`
	Title     string      //`json:"title"`
	Forecasts []forecasts //`json:"forecasts"`
}

type forecasts struct {
	Date   string `json:"date"`
	AAAA   string // スペルミスしてみる
	Telop  string `json:"telop"`
	Detail detail `json:"detail"`
}

type detail struct {
	Weather string `json:"weather"`
	Wind    string `json:"wind"`
	Wave    string `json:"wave"`
}

// type StatusJSON interface{}

func main() {
	http.HandleFunc("/json", handleJSON)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	var statusJson StatusJSON
	json.NewDecoder(r.Body).Decode(&statusJson)
	fmt.Println(statusJson)
	fmt.Fprint(w, "OK")
}
