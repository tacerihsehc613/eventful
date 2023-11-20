package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"rabbit/lib/persistence"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const DB string = "myevents"
const C string = "events"

func main() {
	mongo_url := ""
	if mongo_url = os.Getenv("MONGO_URL"); mongo_url != "" {
		fmt.Println("mongo_url: ", mongo_url)
	}

	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"+":27017"))
	//client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://events-db:27017/events"))
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		panic(err)
	}
	//fmt.Println("client", client)

	coll := client.Database(DB).Collection(C)

	event := persistence.Event{
		ID:        primitive.NewObjectID(),
		Name:      "my festa",
		Duration:  48,       // Duration in minutes, for example
		StartDate: 20231027, // Unix timestamp for the start date
		EndDate:   20231029, // Unix timestamp for the end date
		Location: persistence.Location{
			Name: "ERICA",
		},
	}
	res, err := coll.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}

	fmt.Println("res", res.InsertedID)
	fmt.Println("res", event.ID.Hex())
	fmt.Println("res", event.ID)

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	//results := []bson.M{}
	results := []persistence.Event{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println("results: ", results[0])
	fmt.Printf("type: %T\n", results[0])

	for _, result := range results {
		if !result.ID.IsZero() {
			//if result["_id"] != (primitive.ObjectID{}) {
			//fmt.Printf("Valid _id: %s %T\n", result["_id"], result["_id"])
			fmt.Printf("Valid _id: %s\n", result.ID.Hex())
			fmt.Println(result.ID.IsZero())
			fmt.Println(result.Location.ID.IsZero())
		} else {
			fmt.Println("Invalid _id")
			fmt.Println(result.ID.IsZero())
		}
	}

	fmt.Println("one event retrieval")
	objectId, err := primitive.ObjectIDFromHex("653c957bb27c785dbd19328d")
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", objectId)

	eventIDAsBytes, _ := objectId.MarshalText()
	fmt.Println("eventIDAsBytes", eventIDAsBytes)

	/* filter := bson.M{"_id": objectId}
	var event persistence.Event
	err = coll.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		panic(err)
	}
	fmt.Println("event", event)
	fmt.Printf("type: %T\n", event)

	filter = bson.M{"name": "A Event"}
	var event2 persistence.Event
	err = coll.FindOne(context.TODO(), filter).Decode(&event2)

	if err != nil {
		panic(err)
	}
	fmt.Println("event2", event2)
	fmt.Printf("type: %T\n", event2)

	fmt.Println("loc")
	var location persistence.Location
	fmt.Println("loc", location)

	location2 := persistence.Location{}
	fmt.Println("loc2", location2)
	location.Halls = append(location.Halls, persistence.Hall{Name: "A", Location: "B"})
	//fmt.Println(location.Halls == location2.Halls)
	for i := range location.Halls {
		fmt.Println(i, location.Halls[i])
	} */
	// event = persistence.Event{
	// 	ID:        primitive.NewObjectID(),
	// 	Name:      "AAAAA Event",
	// 	Duration:  48,       // Duration in minutes, for example
	// 	StartDate: 20231027, // Unix timestamp for the start date
	// 	EndDate:   20231029, // Unix timestamp for the end date
	// 	Location: persistence.Location{
	// 		Name: "ERICA",
	// 	},
	// }
	// res, err := coll.InsertOne(context.TODO(), event)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("res", res.InsertedID)
	// fmt.Println("res", event.ID.Hex())
	// fmt.Println("res", event.ID)

}
