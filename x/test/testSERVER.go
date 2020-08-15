package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {

	http.HandleFunc("/hello", helloworld)

	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()

}
