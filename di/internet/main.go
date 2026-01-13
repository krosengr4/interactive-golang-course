package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	fmt.Println("Check it out: http://localhost:5001")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreeterHandler)))
}
