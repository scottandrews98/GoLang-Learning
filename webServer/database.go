package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"strings"
)

// Test this url http://localhost:9090/getpages?request=test&apikey=LqEYJ1rpIntd8A9ThQfwHqrypdhWCUDKc3jUQjr4YGrG21AxUhMMJRVhb8dh

func getPages(w http.ResponseWriter, r *http.Request){
	// Connects to database
	db, err := sql.Open("mysql", ""+ os.Getenv("dbUser") +":"+ os.Getenv("dbPass")+"@tcp("+ os.Getenv("dbHost") +":"+ os.Getenv("dbPort") +")/"+ os.Getenv("dbName") +"?charset=utf8")

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

			fmt.Fprintf(w, strconv.Itoa(totalCount))
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
	rows, err := db.Query("SELECT id FROM users WHERE api_token = "+ apiKey +" LIMIT 1")
	checkErr(err)

	var id int
	var rowCount int

	for rows.Next() {
		
		err = rows.Scan(&id)
		checkErr(err)
		rowCount += 1
	}

	if rowCount != 0{
		return id, false
	}else{
		return 0, true
	}
}

func viewCount(db *sql.DB, userID int) int{
	// Database query	
	rows, err := db.Query("SELECT COUNT(*) as count FROM users WHERE id = "+ strconv.Itoa(userID) +"")
	checkErr(err)

	var count int

	for rows.Next() {
		err = rows.Scan(&count)
		checkErr(err)
	}

	return count
}