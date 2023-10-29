package listener

import (
	"log"
	"rabbit/contracts"
	"rabbit/lib/msgqueue"
	"rabbit/lib/persistence"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventProcessor struct {
	EventListener msgqueue.EventListener
	Database      persistence.DatabaseHandler
}

func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events...")

	received, errors, err := p.EventListener.Listen("eventCreated")

	if err != nil {
		return err
	}

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
		case err := <-errors:
			log.Printf("Received error while consuming msg: %s", err)
		}
	}
}

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)
		id, err := primitive.ObjectIDFromHex(e.ID)
		if err != nil {
			log.Printf("Error parsing ObjectID: %v", err)
			// Handle the error as needed
		} else {
			p.Database.AddEvent(persistence.Event{ID: id})
		}
		//p.Database.AddEvent(persistence.Event{ID: bson.ObjectId(e.ID)})
	case *contracts.LocationCreatedEvent:
		log.Printf("location %s created: %v", e.ID, e)
		id, err := primitive.ObjectIDFromHex(e.ID)
		if err != nil {
			log.Printf("Error parsing ObjectID: %v", err)
			// Handle the error as needed
		} else {
			p.Database.AddLocation(persistence.Location{ID: id})
		}
		//p.Database.AddLocation(persistence.Location{ID: bson.ObjectId(e.ID)})
	default:
		log.Printf("unknown event type: %t", e)
	}
}
