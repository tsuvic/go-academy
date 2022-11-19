package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseJSON struct {
	// Message string
	// Status  int

	Message string `json:"message"`
	Status  int    `json:"status"`
}

func main() {
	http.HandleFunc("/json", handleJSON)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	resJson := ResponseJSON{"JSON返却", 200}
	res, err := json.Marshal(resJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}
