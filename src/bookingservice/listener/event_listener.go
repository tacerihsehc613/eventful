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
	//defer close(received)
	//defer close(errors)

	for {
		select {
		case evt := <-received:
			p.handleEvent(evt)
			// if err := p.handleEvent(evt); err != nil {
			// 	log.Printf("Error processing event: %s", err)
			// }
		case err := <-errors:
			log.Printf("Received error while consuming msg: %s", err)
		}
	}
}

/*func (p *EventProcessor) ProcessEvents() error {
	log.Println("Listening to events...")

	received, errors, err := p.EventListener.Listen("eventCreated")

	log.Println("Test Listen")
	if err != nil {
		return err
	}

	// for {
	// 	select {
	// 	case evt := <-received:
	// 		p.handleEvent(evt)
	// 	case err := <-errors:
	// 		log.Printf("Received error while consuming msg: %s", err)
	// 	}
	// }

	done := make(chan struct{}, 1) // Channel to signal completion
	defer close(done)
	go func() {
		defer close(done)

		for {
			select {
			case evt, ok := <-received:
				if !ok {
					return
				}
				p.handleEvent(evt)
			case err, ok := <-errors:
				if !ok {
					return
				}
				log.Printf("Received error while consuming msg: %s", err)
			}
		}
	}()

	// Wait for the goroutine to finish before returning
	<-done
	return nil

}*/

func (p *EventProcessor) handleEvent(event msgqueue.Event) {
	switch e := event.(type) {
	case *contracts.EventCreatedEvent:
		log.Printf("event %s created: %s", e.ID, e)
		id, err := primitive.ObjectIDFromHex(e.ID)
		id2, err2 := primitive.ObjectIDFromHex(e.LocationID)
		if err != nil || err2 != nil {
			log.Printf("Error parsing ObjectID: %v", err)
			// Handle the error as needed
		} else {
			id, _, err := p.Database.AddEvent4Booking(persistence.Event{ID: id, Location: persistence.Location{ID: id2}})
			log.Printf("id created: %s", id)
			if err != nil {
				log.Printf(`{error: Error occured while persisting event %s}`, err)
				return
			}
		}
		//_, _, err = p.Database.AddEvent4Booking(persistence.Event{ID: id, Location: persistence.Location{ID: id2}})
		// if err != nil {
		// 	return fmt.Errorf("error persisting event: %s", err)
		// }
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
	//return nil
}
