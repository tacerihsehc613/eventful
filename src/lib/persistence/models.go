package persistence

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	First    string
	Last     string
	Age      int
	Bookings []Booking
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %s, First: %s, Last: %s, Age: %d, Bookings: %v", u.ID.Hex(), u.First, u.Last, u.Age, u.Bookings)
}

type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             //`bson:"name"`
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

// omitempty를 할 경우, 빈 값이면 생략된다.
// ObjectID('000000000000000000000000') 조차도 생략된다.
type Location struct {
	//ID        bson.ObjectId `bson:"_id"`
	ID        primitive.ObjectID `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}
