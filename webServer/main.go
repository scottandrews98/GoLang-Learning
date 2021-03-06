package main

import(
	"fmt"
	"net/http"
	"log"
	"html"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
	"strings"
)

// Function that starts the web server 
func main(){
	setEnviroment()

	// Main application routes
	http.HandleFunc("/", home)
	http.HandleFunc("/api/", handleAPIRoutes)

	err := http.ListenAndServe(":9090", nil)

	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/en/03.2.html
// https://tutorialedge.net/golang/writing-a-twitter-bot-golang/

func handleAPIRoutes(w http.ResponseWriter, r *http.Request){
	// Connects to database
	db, err := sql.Open("mysql", ""+ os.Getenv("dbUser") +":"+ os.Getenv("dbPass")+"@tcp("+ os.Getenv("dbHost") +":"+ os.Getenv("dbPort") +")/"+ os.Getenv("dbName") +"?charset=utf8")
	w.Header().Set("Content-Type", "application/json")

	if err != nil{
		checkErr(err)
	}

	apiKey := r.URL.Query()["apikey"]
	websiteID := r.URL.Query()["websiteid"]
	dateFrom := r.URL.Query()["datefrom"]
	dateTo := r.URL.Query()["dateto"]

	if len(apiKey) > 0{
		apiConvert := strings.Join(apiKey, " ")
		websiteIDConvert := strings.Join(websiteID," ")
		dateFromConvert := strings.Join(dateFrom, " ")
		dateToConvert := strings.Join(dateTo, " ")

		var apiReponse []byte
		var invalidAPIKey bool

		// Check api Key First on any route thats /api unless its the fetch websites route
		if html.EscapeString(r.URL.Path) == "/api/getwebsites"{
			invalidAPIKey = false
		}else{
			invalidAPIKey = checkAPIKey(db, apiConvert, websiteIDConvert)
		}

	
		if invalidAPIKey == false{
			switch html.EscapeString(r.URL.Path) {
				case "/api/getviews":
					// Check to make sure values are expected
					valuesExpected := checkRequest([]string{websiteIDConvert})

					if valuesExpected == true{
						apiReponse, err = viewCount(db, websiteIDConvert, dateFromConvert, dateToConvert)
					}else{
						badRequest(w)
					}	
					
				case "/api/getwebsites":
					apiReponse, err = websites(db, apiConvert)

				default:
					fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
			}

			if err == nil{
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, string(apiReponse))
			}else{
				checkErr(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error Processing Request")
			}
		}else{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid API Key")
		}
	}else{
		badRequest(w)
	}
}

// Makes sure api key is valid and exisits in the database
func checkAPIKey(db *sql.DB, apiKey string, websiteID string) bool{
	rows, err := db.Query("SELECT users.id FROM users INNER JOIN websites on users.id = websites.user_id WHERE api_token = '"+ apiKey +"' AND websites.id = "+ websiteID +" LIMIT 1")
	checkErr(err)

	var id int
	var rowCount int

	for rows.Next() {
		err = rows.Scan(&id)
		checkErr(err)
		rowCount += 1
	}

	if rowCount > 0{
		return false
	}else{
		return true
	}
}

// Logs any errors that could occur during processing
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Checks for the correct parameters in each request
func checkRequest(requiredValues []string) bool{
	emptyVariable := true

	for i := 0; i < len(requiredValues); i++ { 
        if len(requiredValues[i]) <= 0{
			emptyVariable = false
		}
    } 
  
	return emptyVariable
}

func badRequest(w http.ResponseWriter){
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Invalid Request Parameters")
}