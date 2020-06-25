package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sessions struct {
	golfType string `bson:"golfType"`
	value    int    `bson:"value"`
}

func getSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//var results = make([]*Sessions, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("golfPlayer").Collection("sessions")

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(ctx) {

		sessions := Sessions{}
		err := cursor.Decode(&sessions)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(cursor)
		}

		//results = append(results, sessions)
	}

	//fmt.Fprintf(w, results)
	// fmt.Println(results)

	defer cancel()

	// pagesJson, err := json.Marshal(results)
	// if err != nil {
	// 	log.Fatal("Cannot encode to JSON ", err)
	// }
	// fmt.Fprintf(os.Stdout, "%s", pagesJson)
}
