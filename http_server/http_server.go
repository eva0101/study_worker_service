package http_server

import (
	"errors"
	"net/http"
	"study_docker/handlers"

	"github.com/gorilla/mux"
)

func StartHTTPServer() error {
	router := mux.NewRouter()

	router.HandleFunc("/employees", handlers.AddEmployeeHandler).Methods("POST")
	router.HandleFunc("/employees", handlers.GetListEmployeesHandler).Methods("GET")
	router.HandleFunc("/employees", handlers.DeleteEmployeHandler).Methods("DELETE")

	err := http.ListenAndServe(":5050", router)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
