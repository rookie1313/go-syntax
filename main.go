package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(MyGreeterHandler))
	if err != nil {
		return
	}
}

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
