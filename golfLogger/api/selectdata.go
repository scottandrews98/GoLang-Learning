package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
