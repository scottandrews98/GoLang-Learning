package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
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
	ID primitive.ObjectID `json:"Id,omitempty" bson:"_id,omitempty"`
	// Driver     int32              `json:"Driver,omitempty" bson:"Driver,omitempty"`
	// ThreeWood  int32              `json:"Three_Wood,omitempty" bson:"Three_Wood,omitempty"`
	// FiveHybrid int32              `json:"Five_Hybrid,omitempty" bson:"Five_Hybrid,omitempty"`
	// SixIron    int32              `json:"Six_Iron,omitempty" bson:"Six_Iron,omitempty"`
	// SevernIron int32              `json:"Severn_Iron,omitempty" bson:"Severn_Iron,omitempty"`
	// EightIron  int32              `json:"Eight_Iron,omitempty" bson:"Eight_Iron,omitempty"`
	// NineIron   int32              `json:"Nine_Iron,omitempty" bson:"Nine_Iron,omitempty"`
	// PW         int32              `json:"PW,omitempty" bson:"PW,omitempty"`
	//SW int32 `json:"SW,omitempty" bson:"SW,omitempty"`
}

// Pulls in the latest club length document
func getClubLengths(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("golfPlayer").Collection("clubs")

	var clubs []Clubs

	err = collection.FindOne(ctx, bson.M{}).Decode(&clubs)
	//result := collection.FindOne(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	// for cursor.Next(ctx) {
	// 	var session Sessions

	// 	err := cursor.Decode(&session)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	sessions = append(sessions, session)
	// }

	defer cancel()

	json.NewEncoder(w).Encode(clubs)
}
