package main

import(
	//"fmt"
	"net/http"
	"log"
)

// Function that starts the web server 
func main(){
	setEnviroment()

	// Main application routes
	http.HandleFunc("/", home)
	http.HandleFunc("/getpages", getPages)

	err := http.ListenAndServe(":9090", nil)

	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html