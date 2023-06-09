package main

import (
	"fmt"
	"os"
)

func main() {
	isPanic := throwsPanic(f)
	fmt.Println(isPanic)
}

func f() {
	var user = os.Getenv("USER")
	if user == "pro" {
		panic("test for panic")
	}
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}
