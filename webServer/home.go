package main

import(
	"fmt"
	"net/http"
	//"log"
)

// Function to handle /
func home(w http.ResponseWriter, r *http.Request){
	fmt.Println("new visitor")
	fmt.Fprintf(w, "Hello Home")
}