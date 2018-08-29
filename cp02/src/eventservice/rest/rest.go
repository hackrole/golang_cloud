package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/golang_cloud/cp02/src/lib/persistence"
	"todo.com/myevents/lib/msgqueue"
)

func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler, eventEmitter emitter.EventEmitter) error {
	handler := NewEventHandler(databasehandler, eventEmitter)
	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()
	eventsrouter.Methods("GET").Path("/{SearchCriterial}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.AllEventhandler)
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
