package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		wellHitShots := r.FormValue("wellHit")
		wellHitInt, err := strconv.Atoi(wellHitShots)

		if err != nil {
			checkError(err)
		}

		inputResponse := newDocument(golfType, shotsInt, wellHitInt)

		if inputResponse == true {
			fmt.Fprintf(w, "Data Added")
		} else {
			fmt.Fprintf(w, "Error Adding Data")
		}
	default:
		fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
	}
}

func newDocument(golfType string, shots int, wellHit int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("databaseURL")))

	collection := client.Database("golfPlayer").Collection("sessions")
	collection.InsertOne(ctx, bson.M{"golfType": golfType, "value": shots, "totalWellHit": wellHit})

	if err != nil {
		checkError(err)
	}

	defer cancel()

	return true
}

func updateDocument(id string, distance int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("databaseURL")))

	collection := client.Database("golfPlayer").Collection("clubs")

	convertedID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": convertedID}
	change := bson.M{
		"$set": bson.M{
			"clubDistance": distance,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, change)

	defer cancel()

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func checkError(err error) {
	fmt.Printf("err: %v", err)
}

func updateDistance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Must Be A Post Request To This Endpoint")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		distance := r.FormValue("distance")
		documentID := r.FormValue("id")
		distanceInt, err := strconv.Atoi(distance)

		if err != nil {
			checkError(err)
		}

		inputResponse := updateDocument(documentID, distanceInt)

		if inputResponse == true {
			fmt.Fprintf(w, "Data Updated")
		} else {
			fmt.Fprintf(w, "Error Updating Data")
		}
	default:
		fmt.Fprintf(w, "Welcome To The No Track Website Stats API")
	}
}

type GolfCourses struct {
	Name   string  `json:"Name,omitempty"`
	Rating float64 `json:"Rating,omitempty"`
}

// Finds nearest golf courses depending on postcode entered
func findGolf(w http.ResponseWriter, r *http.Request) {
	apiKey := os.Getenv("googleAPIKey")

	resp, err := http.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=driving+range+in+Crewe&key=" + apiKey + "")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	// https://github.com/tidwall/gjson A great way of unmarshlling json if your use to javascript
	result := gjson.GetBytes(data, "results")

	var courses []GolfCourses

	result.ForEach(func(key, value gjson.Result) bool {
		SiteName := value.Get("name")
		SiteRating := value.Get("rating")

		course := GolfCourses{SiteName.String(), SiteRating.Float()}
		courses = append(courses, course)

		return true // keep iterating
	})

	json.NewEncoder(w).Encode(courses)
}
