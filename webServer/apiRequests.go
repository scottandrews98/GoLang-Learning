package main

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
	"time"
)

func viewCount(db *sql.DB, websiteID string, dateFrom string, dateTo string) ([]byte, error){
	type views struct{
		WebsiteURL string `json:"WebsiteURL"`
		WebsiteViews int `json:"WebsiteViews"`
	}

	var rows *sql.Rows
	var err error

	if len(dateFrom) == 0 && len(dateTo) == 0{
		rows, err = db.Query("SELECT count(views.id) as views, websites.websiteURL as websiteURL FROM views INNER JOIN websites ON views.website_id = websites.id WHERE views.website_id = "+ websiteID +"")
	}else{
		// User has specified dates so change query to include this
		dateFromFormated, err := time.Parse(layoutISO, date)

		rows, err = db.Query("SELECT count(views.id) as views, websites.websiteURL as websiteURL FROM views INNER JOIN websites ON views.website_id = websites.id WHERE views.website_id = "+ websiteID +" AND created_at BETWEEN "'+ $formattedDateFrom +' 00:00:00" AND "'+ $formattedDateTo +' 23:59:59"")
	}

	
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

func websites(db *sql.DB, apiKey string) ([]byte, error){
	type websites struct{
		WebsiteID int `json:"WebsiteID"`
		WebsiteURL string `json:"WebsiteURL"`
		AddedTrackingCode bool `json:"AddingTrackingCode"`
	}

	var active int

	// Database query	
	rows, err := db.Query("SELECT websites.id as websiteID, websites.websiteURL as websiteURL, websites.addedTrackingCode as trackingCode FROM websites INNER JOIN users ON websites.user_id = users.id WHERE users.api_token = '"+ apiKey +"'")
	checkErr(err)

	websiteArray := make([]*websites, 0)

	for rows.Next() {
		website := new(websites)
		err = rows.Scan(&website.WebsiteID, &website.WebsiteURL, &active)
		checkErr(err)

		if active == 1{
			website.AddedTrackingCode = true
		}else{
			website.AddedTrackingCode = false
		}

		websiteArray = append(websiteArray, website)
	}

	return json.Marshal(websiteArray)
}