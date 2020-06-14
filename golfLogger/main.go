package main

// import (
//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )


// Main job is to log to a mongo db database about each golf session
import(
	"fmt"
	"net/http"
	"log"
)

// Function that starts the web server 
func main(){
	
	// Main application routes
	//http.HandleFunc("/", home)
	http.HandleFunc("/api/", apiRequest)

	err := http.ListenAndServe(":9090", nil)

	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}

func apiRequest(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
}