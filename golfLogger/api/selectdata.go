package main

import (
	"context"
	"encoding/json"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sessions struct {
	//ID       primitive.ObjectID `json:"Id,omitempty" bson:"_id,omitempty"`
	GolfType string `json:"GolfType,omitempty" bson:"golfType,omitempty"`
	Value    int    `json:"Value,omitempty" bson:"value,omitempty"`
}

func getSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("databaseURL")))
	collection := client.Database("golfPlayer").Collection("sessions")

	cursor, err := collection.Find(ctx, bson.M{})
	var sessions []Sessions

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var session Sessions

		err := cursor.Decode(&session)
		if err != nil {
			log.Fatal(err)
		}
		sessions = append(sessions, session)
	}

	defer cancel()

	json.NewEncoder(w).Encode(sessions)
}

type Clubs struct {
	ID           primitive.ObjectID `json:"Id,omitempty" bson:"_id,omitempty"`
	ClubName     string             `json:"clubName,omitempty" bson:"clubName,omitempty"`
	ClubDistance int32              `json:"clubDistance,omitempty" bson:"clubDistance,omitempty"`
}

// Pulls in the latest club length document
func getClubLengths(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("databaseURL")))
	collection := client.Database("golfPlayer").Collection("clubs")

	var clubs []Clubs

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {
		var club Clubs

		err := cursor.Decode(&club)
		if err != nil {
			log.Fatal(err)
		}
		clubs = append(clubs, club)
	}

	defer cancel()

	json.NewEncoder(w).Encode(clubs)
}

type ShotCount struct {
	ID         string `bson:"_id,omitempty"`
	TotalShots int    `bson:"TotalGolfShots,omitempty,truncate"`
}

func getShotsAndAverages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("databaseURL")))
	collection := client.Database("golfPlayer").Collection("sessions")
	var aggregateExpression primitive.D

	if html.EscapeString(r.URL.Path) == "/api/getshots" {
		aggregateExpression = bson.D{{"$group", bson.M{"_id": "null", "TotalGolfShots": bson.M{"$sum": "$value"}}}}
	} else {
		aggregateExpression = bson.D{{"$group", bson.M{"_id": "null", "TotalGolfShots": bson.M{"$avg": "$totalWellHit"}}}}
	}

	pipe, err := collection.Aggregate(ctx, mongo.Pipeline{aggregateExpression})

	var loadedStruct []ShotCount

	if err = pipe.All(ctx, &loadedStruct); err != nil {
		panic(err)
	}

	defer cancel()

	json.NewEncoder(w).Encode(loadedStruct)
}
