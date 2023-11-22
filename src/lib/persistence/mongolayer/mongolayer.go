package mongolayer

import (
	"context"
	"log"
	"rabbit/lib/persistence"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB        = "myevents"
	USERS     = "users"
	EVENTS    = "events"
	LOCATIONS = "locations"
)

type MongoDBLayer struct {
	//session *mgo.Session
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	/* s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		session: s,
	}, err */
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		client: client,
	}, nil
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) (string, error) {
	//s := mgoLayer.getFreshSession()
	//defer s.Close()
	coll := mgoLayer.client.Database(DB).Collection(EVENTS)

	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}
	if e.Location.ID.IsZero() {
		e.Location.ID = primitive.NewObjectID()
	}
	//return e.ID.Hex(), s.DB(DB).C(EVENTS).Insert(e)
	_, err := coll.InsertOne(context.TODO(), e)
	if err != nil {
		return "", err
	}

	return e.ID.Hex(), nil

}

func (mgoLayer *MongoDBLayer) AddLocation(l persistence.Location) (persistence.Location, error) {
	//s := mgoLayer.getFreshSession()
	//defer s.Close()
	//l.ID = bson.NewObjectId()
	l.ID = primitive.NewObjectID()
	coll := mgoLayer.client.Database(DB).Collection(LOCATIONS)
	_, err := coll.InsertOne(context.TODO(), l)
	if err != nil {
		return l, err
	}
	//err := s.DB(DB).C(LOCATIONS).Insert(l)
	return l, err
}

/* func (mgoLayer *MongoDBLayer) AddBookingForUser(id []byte, b persistence.Booking) error {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$addToSet": bson.M{"bookings": b}})
} */

func (mgoLayer *MongoDBLayer) AddBookingForUser(id string, b persistence.Booking) error {
	coll := mgoLayer.client.Database(DB).Collection(USERS)
	filter := bson.M{"_id": id}

	//Define the update operation.
	update := bson.M{
		"$addToSet": bson.M{"bookings": b},
	}
	opts := options.Update().SetUpsert(true)
	// Use the UpdateOne method to apply the update.
	_, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

/* func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	//e := persistence.Event{}
	var e persistence.Event
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}*/

func (mgoLayer *MongoDBLayer) FindEvent(id string) (persistence.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return persistence.Event{}, err
	}

	coll := mgoLayer.client.Database(DB).Collection(EVENTS)
	filter := bson.M{"_id": objectId}
	var event persistence.Event
	err = coll.FindOne(context.TODO(), filter).Decode(&event)
	return event, err
}

func (mgoLayer *MongoDBLayer) FindLocation(id string) (persistence.Location, error) {
	/* s := mgoLayer.getFreshSession()
	defer s.Close()
	location := persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(bson.M{"_id": bson.ObjectId(id)}).One(&location)
	return location, err */
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return persistence.Location{}, err
	}

	coll := mgoLayer.client.Database(DB).Collection(LOCATIONS)
	filter := bson.M{"_id": objectId}
	//var location persistence.Location
	location := persistence.Location{}
	err = coll.FindOne(context.TODO(), filter).Decode(&location)
	return location, err
}

/*
	func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
		s := mgoLayer.getFreshSession()
		defer s.Close()
		e := persistence.Event{}
		err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
		return e, err
	}
*/
func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	coll := mgoLayer.client.Database(DB).Collection(EVENTS)
	filter := bson.M{"name": name}
	var event persistence.Event
	err := coll.FindOne(context.TODO(), filter).Decode(&event)
	return event, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	//s := mgoLayer.getFreshSession()
	//defer s.Close()
	coll := mgoLayer.client.Database(DB).Collection(EVENTS)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	results := []persistence.Event{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
	//var events []persistence.Event
	//err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	//return events, err
}

func (mgoLayer *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	//s := mgoLayer.getFreshSession()
	//defer s.Close()
	coll := mgoLayer.client.Database(DB).Collection(LOCATIONS)
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	results := []persistence.Location{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil

	//var locations []persistence.Location
	//err := s.DB(DB).C(LOCATIONS).Find(nil).All(&locations)
	//return locations, err
}

//The new driver abstracts session management
/* func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
} */
