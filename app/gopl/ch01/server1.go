// server1 は最小の echo サーバ
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 個々のリクエストに対して handler が呼ばれる
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler は、リクエストされた URL r の Path 要素を返す
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

