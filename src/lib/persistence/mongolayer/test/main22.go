package main

/* const DB string = "myevents"
const C string = "events"

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	//fmt.Println("client", client)

	coll := client.Database(DB).Collection(C)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	results := []bson.M{}
	//results := []persistence.Event{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println("results: ", results[0])
	fmt.Printf("type: %T\n", results[0])

	for _, result := range results {
		if id, ok := result["_id"].(primitive.ObjectID); ok && !id.IsZero() {
			//if result["_id"] != (primitive.ObjectID{}) {
			//fmt.Printf("Valid _id: %s %T\n", result["_id"], result["_id"])
			fmt.Printf("Valid _id: %s\n", id.Hex())
			fmt.Println(id.IsZero())
		} else {
			fmt.Println("Invalid _id")
			fmt.Println(id.IsZero())
		}
	}

	//

	event := persistence.Event{
		//ID: primitive.NewObjectID(),
		//ID:        primitive.ObjectID{},
		Name:      "XXX Event",
		Duration:  48,       // Duration in minutes, for example
		StartDate: 20231027, // Unix timestamp for the start date
		EndDate:   20231029, // Unix timestamp for the end date
		Location: persistence.Location{
			//ID:   primitive.NewObjectID(),
			//ID:   locationID,
			ID:   primitive.ObjectID{},
			Name: "island",
		},
	}
	res, err := coll.InsertOne(context.TODO(), event)
	if err != nil {
		panic(err)
	}

	fmt.Println("res", res.InsertedID)
	fmt.Println("res", event.ID.Hex())
	fmt.Println("res", event.ID)
	fmt.Println("res loc", event.Location.ID)

} */
