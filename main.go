package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":3434", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello this BLUE!!!")
}