package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/location", handleLocation)
	http.HandleFunc("/search", handleSearch)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

// URLのクエリを出力
func handleLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")

	fmt.Printf("your queries are %s %s \n", name, category)
}

// postデータのparseと出力
func handleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.Form)
	fmt.Println(r.Body)

	r.ParseForm()
	form := r.PostForm
	fmt.Println(form)

	keyword, ok := form["keyword"]
	if !ok {
		fmt.Println("no keyword")
	}

	category, ok := form["category"]
	if !ok {
		fmt.Println("no category")
	}

	fmt.Printf("your queries are %s %s \n", keyword, category)

}
