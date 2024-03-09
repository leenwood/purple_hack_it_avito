package main

import "net/http"

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8998", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

}
