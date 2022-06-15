package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, word string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", word)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func MyGreeterHandler(w http.ResponseWriter, _ *http.Request) {
	Greet(w, "leo")
}

func main() {
	err := http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
