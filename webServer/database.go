package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"strings"
	"html"
)

// Test this url http://localhost:9090/getpages?request=test&apikey=LqEYJ1rpIntd8A9ThQfwHqrypdhWCUDKc3jUQjr4YGrG21AxUhMMJRVhb8dh

func getPages(w http.ResponseWriter, r *http.Request){
	// Connects to database
	db, err := sql.Open("mysql", ""+ os.Getenv("dbUser") +":"+ os.Getenv("dbPass")+"@tcp("+ os.Getenv("dbHost") +":"+ os.Getenv("dbPort") +")/"+ os.Getenv("dbName") +"?charset=utf8")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err != nil{
		checkErr(err)
	}else{
		fmt.Println("database connected")
	}

	//fmt.Println("Get params are", r.URL.Query())
	apiKey := r.URL.Query()["apikey"]
	dataRequest := r.URL.Query()["request"]
	
	if len(apiKey) > 0 && len(dataRequest) > 0{
		//var apiConvert string
		apiConvert := strings.Join(apiKey," ")

		userID, invalidAPIKey := checkAPIKey(db, apiConvert)

		if invalidAPIKey == false{
			totalCount := viewCount(db, userID)
			fmt.Printf(totalCount)
			fmt.Fprintf(w, "Total Page Views: " +html.UnescapeString(totalCount)+ "")
		}
	}else{
		fmt.Fprintf(w, "Invalid Params")
	}



}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Makes sure api key is valid and exisits in the database
func checkAPIKey(db *sql.DB, apiKey string) (int, bool){
	rows, err := db.Query("SELECT id FROM users WHERE api_token = '"+ apiKey +"' LIMIT 1")
	checkErr(err)

	var id int
	var rowCount int

	for rows.Next() {
		
		err = rows.Scan(&id)
		checkErr(err)
		rowCount += 1
	}

	if rowCount > 0{
		return id, false
	}else{
		return 0, true
	}
}

func viewCount(db *sql.DB, userID int) string{
	type websites struct{
		id int
		websiteURL string
	}

	// Database query	
	rows, err := db.Query("SELECT id, websiteURL FROM websites WHERE user_id = "+ strconv.Itoa(userID) +"")
	checkErr(err)

	websiteArray := make([]*websites, 0)
	var returnSelect string

	for rows.Next() {
		website := new(websites)

		err = rows.Scan(&website.id, &website.websiteURL)
		checkErr(err)

		websiteArray = append(websiteArray, website)
	}

	// Loops through all websites and returns a dropdown
	if len(websiteArray) > 1{
		returnSelect = "<select>"

		for i := 0; i < len(websiteArray); i++ {
			returnSelect+= "<option>"+websiteArray[i].websiteURL+"</option>"
		}

		returnSelect += "</select>"
	}else{	
		returnSelect = "<select><option>"+websiteArray[0].websiteURL+"</option></select>"
	}	

	return returnSelect
}