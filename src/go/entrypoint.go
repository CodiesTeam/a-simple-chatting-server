package main

import (
	"fmt"
	"log"
	"net/http"
	"a-simple-chatting-server/src/go/db"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	err := http.ListenAndServe("0.0.0.0:8091", h)
	if err != nil {
		log.Fatal(err)
	}
	db.CreateDB("testdb")
}