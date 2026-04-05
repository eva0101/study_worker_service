package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"study_docker/employee"
	"study_docker/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AddEmployeeHandler(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	httpBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err read body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var employe employee.Employee
	if err = json.Unmarshal(httpBody, &employe); err != nil {
		fmt.Println("err unmarshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sql.InsertEmployee(r.Context(), pool, employe)

	w.Write([]byte("новый работник был добавлен"))
}

func GetListEmployeesHandler(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	list := sql.SelectEmployees(r.Context(), pool)

	jsonList, err := json.Marshal(list)
	if err != nil {
		fmt.Println("err marshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonList)
}

func DeleteEmployeHandler(pool *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	httpBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err read body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var id employee.Employee
	if err = json.Unmarshal(httpBody, &id); err != nil {
		fmt.Println("err unmarshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	notification := sql.DeleteEmployee(r.Context(), pool, id)
	if notification == "" {
		w.Write([]byte("сотрудник был удален"))
	} else {
		w.Write([]byte(notification))
	}
}
