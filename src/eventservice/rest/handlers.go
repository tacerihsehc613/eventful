package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rabbit/contracts"
	"rabbit/lib/msgqueue"
	"rabbit/lib/persistence"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type eventServiceHandler struct {
	dbhandler    persistence.DatabaseHandler
	eventEmitter msgqueue.EventEmitter
}

func NewEventHandler(databasehandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) *eventServiceHandler {
	return &eventServiceHandler{
		dbhandler:    databasehandler,
		eventEmitter: eventEmitter,
	}
}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteria, ok := vars["SearchCriteria"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: "Missing Search Criteria, you can either
		search by id via /id/4 
		to search by name via /name/coldplayconcert"}`)
		return
	}
	searchkey, ok := vars["search"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{error: No search keys found, you can either
		search by id via /id/4 
		to search by name via /name/coldplayconcert"}`)
		return
	}
	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
		/* id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbhandler.FindEvent(id)
		} */
		event, err = eh.dbhandler.FindEvent(searchkey)
	}
	if err != nil {
		fmt.Fprintf(w, "{error %s}", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbhandler.FindAllAvailableEvents()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Error occured while trying to find all 
			available events %s}`, err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Error occured while trying encode events to JSON %s}`, err)
	}
}

func (eh *eventServiceHandler) oneEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["eventID"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "missing route parameter 'eventID")
		return
	}
	event, err := eh.dbhandler.FindEvent(id)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "event with id %s was not found", id)
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(w).Encode(&event)
}

func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Error occured while decoding event data %s}`, err)
		return
	}
	id, err := eh.dbhandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{error: Error occured while persisting event %s}`, err)
		return
	}
	msg := contracts.EventCreatedEvent{
		//ID:         hex.EncodeToString(id),
		ID:         id,
		Name:       event.Name,
		LocationID: event.Location.ID.Hex(),
		Start:      time.Unix(event.StartDate, 0),
		End:        time.Unix(event.EndDate, 0),
	}
	eh.eventEmitter.Emit(&msg)
	fmt.Fprintf(w, `{id: %s}`, id)

}
