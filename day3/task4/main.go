// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type StructJSON struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// func main() {
// 	http.HandleFunc("/", handleQuery)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal("起動失敗", err)
// 	}
// }

// func handleQuery(w http.ResponseWriter, r *http.Request) {

// 	fmt.Println(r.URL.Query())

// 	name := r.URL.Query().Get("name")
// 	price := r.URL.Query().Get("price")

// 	fmt.Println(name, price)

// 	fmt.Fprint(w, "OK")

// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/form", handleForm)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("起動失敗", err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()

	form := r.PostForm

	key1 := form["key1"]
	key2 := form["key2"]
	key3, ok := form["key3"]

	if !ok {
		fmt.Println("key3は設定なし")
	}

	fmt.Println(key1, key2, key3)

	fmt.Fprint(w, "ok")

}
