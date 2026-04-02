package http_server

import (
	"errors"
	"net/http"
	"study_docker/create_pool"
	"study_docker/handlers"

	"github.com/gorilla/mux"
)

func StartHTTPServer() error {
	router := mux.NewRouter()

	pool := create_pool.CreatePool()
	defer pool.Close()

	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddEmployeeHandler(pool, w, r)
	}).Methods("POST")
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetListEmployeesHandler(pool, w, r)
	}).Methods("GET")
	router.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteEmployeHandler(pool, w, r)
	}).Methods("DELETE")

	err := http.ListenAndServe(":5050", router)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}
