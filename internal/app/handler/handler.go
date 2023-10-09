package handler

import (
	"bytes"
	"employee-service/internal/app/model"
	"employee-service/internal/app/usecase"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	EU *usecase.EmployeeUsecase
}

func NewEmployeeHandler(eu *usecase.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		EU: eu,
	}
}

func (eh *EmployeeHandler) HireEmployee(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Received %s request at %s with body: %s", r.Method, r.URL.Path, string(bodyBytes))

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var e model.Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := eh.EU.CreateEmployee(e); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (eh *EmployeeHandler) FireEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s with query: %s", r.Method, r.URL.Path, r.URL.RawQuery)
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := eh.EU.DeleteEmployee(id); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (eh *EmployeeHandler) GetVacationDays(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s with query: %s", r.Method, r.URL.Path, r.URL.RawQuery)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	days, err := eh.EU.GetEmployeeVacationDays(id)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(days)
	w.Write(jsonData)
}

func (eh *EmployeeHandler) FindEmployee(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s with query: %s", r.Method, r.URL.Path, r.URL.RawQuery)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	employees, err := eh.EU.FindEmployeeByName(name)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(employees)
	w.Write(jsonData)
}
