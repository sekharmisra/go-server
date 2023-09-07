package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandlerFunc)
	http.HandleFunc("/hello", helloHandlerFunc)

	fmt.Println("Server starting at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func formHandlerFunc(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() eror %v\n", err)
	}

	fmt.Fprintf(w, "POST form is successful!")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name is %v\n", name)
	fmt.Fprintf(w, "Address is %v\n", address)
}

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {

	//If wrong path
	if r.URL.Path != "/hello" {
		http.Error(w, "method not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}
