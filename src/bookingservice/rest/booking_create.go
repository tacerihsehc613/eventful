package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"rabbit/contracts"
	"rabbit/lib/msgqueue"
	"rabbit/lib/persistence"
	"time"

	"github.com/gorilla/mux"
)

type eventRef struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type createBookingRequest struct {
	Seats int `json:"seats"`
}

type createBookingResponse struct {
	ID    string   `json:"id"`
	Event eventRef `json:"event"`
}

type CreateBookingHandler struct {
	eventEmitter msgqueue.EventEmitter
	database     persistence.DatabaseHandler
}

func (h *CreateBookingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routeVars := mux.Vars(r)
	var eventID string
	var ok bool
	eventID, ok = routeVars["eventID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "missing route variable 'eventID'")
		return
	}

	eventIDMongo, _ := hex.DecodeString(eventID)
	event, err := h.database.FindEvent(eventIDMongo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error finding event %s", err)
		return
	}

	var bookingRequest createBookingRequest = createBookingRequest{}
	err = json.NewDecoder(r.Body).Decode(&bookingRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not decode JSON body: %s", err)
		return
	}

	if bookingRequest.Seats <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "seats must be a positive number")
		return
	}
	eventIDAsBytes, _ := event.ID.MarshalText()
	booking := persistence.Booking{
		Date:    time.Now().Unix(),
		EventID: eventIDAsBytes,
		Seats:   bookingRequest.Seats,
	}

	msg := contracts.EventBookedEvent{
		EventID: event.ID.Hex(),
		UserID:  "someUserID",
	}

	h.eventEmitter.Emit(&msg)

	h.database.AddBookingForUser([]byte("someUserID"), booking)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&booking)
}
