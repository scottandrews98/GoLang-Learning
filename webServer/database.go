package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	//"strconv"
	"encoding/json"
)

// Great website tutorials on go https://tutorialedge.net/golang/writing-a-twitter-bot-golang/

// Test this url http://localhost:9090/getviews?apikey=sBWauNebh89emzyugEtcwASvEeJzYFljLRwNnaaJ8s3mQ6Ujlr5gFsRCGsS1&websiteid=1

// func getViews(w http.ResponseWriter, r *http.Request){
// 	// Connects to database
// 	db, err := sql.Open("mysql", ""+ os.Getenv("dbUser") +":"+ os.Getenv("dbPass")+"@tcp("+ os.Getenv("dbHost") +":"+ os.Getenv("dbPort") +")/"+ os.Getenv("dbName") +"?charset=utf8")
// 	w.Header().Set("Content-Type", "application/json")

// 	if err != nil{
// 		checkErr(err)
// 	}else{
// 		fmt.Println("database connected")
// 	}

// 	apiKey := r.URL.Query()["apikey"]
// 	websiteID := r.URL.Query()["websiteid"]
	
// 	if len(apiKey) > 0 && len(websiteID) > 0{
// 		apiConvert := strings.Join(apiKey, " ")
// 		websiteIDConvert := strings.Join(websiteID," ")

// 		invalidAPIKey := checkAPIKey(db, apiConvert)

// 		if invalidAPIKey == false{
// 			websiteViews, err := viewCount(db, websiteIDConvert)

// 			if(err == nil){
// 				fmt.Fprintf(w, string(websiteViews))
// 			}else{
// 				fmt.Println(err)
// 			}
			
// 		}else{
// 			fmt.Fprintf(w, "Invalid API Key")
// 		}
// 	}else{
// 		fmt.Fprintf(w, "Invalid Request Parameters")
// 	}
// }




func viewCount(db *sql.DB, websiteID string) ([]byte, error){
	type views struct{
		WebsiteURL string `json:"WebsiteURL"`
		WebsiteViews int `json:"WebsiteViews"`
	}

	// Database query	
	rows, err := db.Query("SELECT count(views.id) as views, websites.websiteURL as websiteURL FROM views INNER JOIN websites ON views.website_id = websites.id WHERE views.website_id = "+ websiteID +"")
	checkErr(err)

	viewArray := make([]*views, 0)

	for rows.Next() {
		view := new(views)
		err = rows.Scan(&view.WebsiteViews, &view.WebsiteURL)
		checkErr(err)
		viewArray = append(viewArray, view)
	}

	return json.Marshal(viewArray)
}