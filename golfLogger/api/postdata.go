package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func apiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		golfType := r.FormValue("golftype")
		shots := r.FormValue("shots")
		//fmt.Fprintf(w, "Name = %s\n", name)
		//fmt.Fprintf(w, "Address = %s\n", address)
		newDocument(golfType, strconv.Atoi(shots))
	default:
		fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
	}
}

func newDocument(golfType string, shots int) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return true
}
