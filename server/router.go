package server

import (
	"net/http"
	"processlist/control"

	"github.com/gorilla/mux"
)

func createHandler() (handler *mux.Router) {

	// creats router
	handler = mux.NewRouter()

	handler.HandleFunc("/", control.HelloWorld).Methods(http.MethodGet)
	handler.HandleFunc("/add", control.AddProcess).Methods(http.MethodPost)
	handler.HandleFunc("/process", control.SendProcess).Methods(http.MethodGet)

	return

}
