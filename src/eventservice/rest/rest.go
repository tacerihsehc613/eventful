package rest

import (
	"net/http"
	"rabbit/lib/msgqueue"
	"rabbit/lib/persistence"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ServeAPI(endpoint string, tlsendpoint string, dbHandler persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) (chan error, chan error) {
	//handler := &eventServiceHandler{dbHandler}
	handler := NewEventHandler(dbHandler, eventEmitter)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("GET").Path("/{eventID}").HandlerFunc(handler.oneEventHandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	//return http.ListenAndServe(":8181", r)
	//return http.ListenAndServe(endpoint, r)

	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	server := handlers.CORS()(r)
	go func() {
		//httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "./cert.pem", "./key.pem", server)
		httptlsErrChan <- http.ListenAndServeTLS(tlsendpoint, "./eventservice/cert.pem", "./eventservice/key.pem", server)
	}()

	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, server)
	}()

	return httpErrChan, httptlsErrChan

}
