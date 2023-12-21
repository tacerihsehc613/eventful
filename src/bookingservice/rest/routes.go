package rest

import (
	"net/http"
	"rabbit/lib/msgqueue"
	"rabbit/lib/persistence"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ServeAPI(listenAddr string, database persistence.DatabaseHandler, eventEmitter msgqueue.EventEmitter) {
	r := mux.NewRouter()
	//r.Methods("post").Path("/events/{eventID}/bookings").Handler(&CreateBookingHandler{eventEmitter: eventEmitter, database: database})
	r.Methods("post").Path("/bookings/{eventID}").Handler(&CreateBookingHandler{eventEmitter: eventEmitter, database: database})

	srv := http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         listenAddr,
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	srv.ListenAndServe()
}
