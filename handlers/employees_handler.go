package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"study_docker/employee"
)

var dbEmpoyees = &employee.ListEmployees{}

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
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

	dbEmpoyees.Add(employe)

	w.Write([]byte("новый работник был добавлен"))
}

func GetListEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	list := dbEmpoyees.Get()

	jsonList, err := json.Marshal(list)
	if err != nil {
		fmt.Println("err marshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonList)
}

func DeleteEmployeHandler(w http.ResponseWriter, r *http.Request) {
	httpBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err read body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var id employee.Num
	if err = json.Unmarshal(httpBody, &id); err != nil {
		fmt.Println("err unmarshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = dbEmpoyees.Delete(id); err != nil {
		w.Write([]byte("такого id нету"))
		fmt.Println(err)
		return
	}

	w.Write([]byte("сотрудник был удален"))
}
