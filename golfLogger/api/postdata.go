package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
		shotsInt, err := strconv.Atoi(shots)

		if err != nil {
			checkError(err)
		}

		inputResponse := newDocument(golfType, shotsInt)

		if inputResponse == true {
			fmt.Fprintf(w, "Data Added")
		} else {
			fmt.Fprintf(w, "Error Adding Data")
		}
	default:
		fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
	}
}

func newDocument(golfType string, shots int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("golfPlayer").Collection("sessions")
	collection.InsertOne(ctx, bson.M{"golfType": golfType, "value": shots})

	if err != nil {
		checkError(err)
	}

	defer cancel()

	return true
}

func checkError(err error) {
	fmt.Printf("err: %v", err)
}
