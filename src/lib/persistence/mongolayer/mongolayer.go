package mongolayer

import (
	"rabbit/lib/persistence"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB        = "myevents"
	USERS     = "users"
	EVENTS    = "events"
	LOCATIONS = "locations"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}
	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)

}

func (mgoLayer *MongoDBLayer) AddLocation(l persistence.Location) (persistence.Location, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	l.ID = bson.NewObjectId()
	err := s.DB(DB).C(LOCATIONS).Insert(l)
	return l, err
}

func (mgoLayer *MongoDBLayer) AddBookingForUser(id []byte, b persistence.Booking) error {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	return s.DB(DB).C(USERS).UpdateId(bson.ObjectId(id), bson.M{"$addToSet": bson.M{"bookings": b}})
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	//e := persistence.Event{}
	var e persistence.Event
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindLocation(id string) (persistence.Location, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	location := persistence.Location{}
	err := s.DB(DB).C(LOCATIONS).Find(bson.M{"_id": bson.ObjectId(id)}).One(&location)
	return location, err
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	var events []persistence.Event
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (mgoLayer *MongoDBLayer) FindAllLocations() ([]persistence.Location, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	var locations []persistence.Location
	err := s.DB(DB).C(LOCATIONS).Find(nil).All(&locations)
	return locations, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
