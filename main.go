package main

import (
	"fmt"
	"net/http"
	"strings"
)

// handler function()
func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                         // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)                   // print form information in server side
	fmt.Println("path :", r.URL.Path)     // print request path
	fmt.Println("scheme: ", r.URL.Scheme) // print request scheme
	fmt.Println(r.Form["url_long"])       // print form information

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello test!")
}

// func CommandLineHandler(writer http.ResponseWriter, request *http.Request) {
// 	url := request.URL
// 	query := url.Query()

// 	id := query["id"]
// 	for i := 0; i < len(id); i++ {
// 		log.Printf("id is %s\n", id[i])
// 	}

// 	name := query.Get("name")
// 	log.Println("first name is", name)
// }

func main() {
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/urlinfo", testHandler)

	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	http.ListenAndServe("localhost:8080", nil)
}
