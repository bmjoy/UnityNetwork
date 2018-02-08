package main

import (
	"log"
	"net/http"
)

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello!"))
}

func main() {
	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
