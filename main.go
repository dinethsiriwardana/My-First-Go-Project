package main

import (
	"fmt"
	"log"
	"net/http"
)

// Every API rout has Response(w) and Request(r)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
		return
	}

	fmt.Fprintf(w, "POST request Successful!\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "address =%s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //Tell Go language to Check the static Dir
	http.Handle("/", fileServer)                        //Handing my root route   / for Index.html
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler) // Just print hello to the screen

	fmt.Printf("Starting server at P 8000\n")

	// Create the server and this is the heart of the program
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
