package main

// import (
//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )

// Main job is to log to a mongo db database about each golf session
import (
	"log"
	"net/http"
)

// Function that starts the web server
func main() {

	// Main application routes
	//http.HandleFunc("/", home)
	http.HandleFunc("/api/addsession", apiRequest)
	http.HandleFunc("/api/getsessions", getSessions)
	http.HandleFunc("/api/getclubs", getClubLengths)
	http.HandleFunc("/api/savedistance", updateDistance)
	http.HandleFunc("/api/getshots", getShotsAndAverages)
	http.HandleFunc("/api/getgoodshots", getShotsAndAverages)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
