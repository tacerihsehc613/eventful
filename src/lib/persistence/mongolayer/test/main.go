package main

const database string = "myevents"
const collection string = "events"

/*func main() {
	// url := "mongodb://localhost:27017"
	// session, _ := mgo.Dial(url)
	// var events []persistence.Event
	// err := session.DB(database).C(collection).Find(nil).All(&events)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ee", events)

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	//fmt.Println("client", client)

	coll := client.Database(database).Collection(collection)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println("results: ", results)

	// _ = client.Ping(ctx, readpref.Primary())
	// var events []persistence.Event
	// collection := client.Database(database).Collection(collection)
	// cursor, err := collection.Find(ctx, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer cursor.Close(ctx)

	// for cursor.Next(ctx) {
	// 	var event persistence.Event
	// 	if err := cursor.Decode(&event); err != nil {
	// 		panic(err)
	// 	}
	// 	events = append(events, event)
	// }
	// if err := cursor.Err(); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("ee", events)

}*/
