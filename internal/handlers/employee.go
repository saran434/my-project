package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"employeejwt/internal/database"
	"employeejwt/internal/models"
	"employeejwt/internal/repository"
)

var store *repository.PostgresStore

func getStore() *repository.PostgresStore {
	if store == nil {
		store = repository.NewPostgresStore(database.DB)
	}
	return store
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		log.Printf("create employee: invalid request body for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := getStore().Create(&emp); err != nil {
		log.Printf("create employee: failed for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Employee Created")
}

func ListEmployees(w http.ResponseWriter, r *http.Request) {
	ems, err := getStore().List()
	if err != nil {
		log.Printf("list employees: failed for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ems)
	fmt.Fprintln(w, "Employees Listed")
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/employees/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("get employee: invalid id for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	e, err := getStore().Get(id)
	if err != nil {
		log.Printf("get employee: failed for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(e)
	fmt.Fprintln(w, "Employee Retrieved")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/employees/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("update employee: invalid id for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		log.Printf("update employee: invalid request body for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := getStore().Update(id, &emp); err != nil {
		log.Printf("update employee: failed for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Employee Updated")
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/employees/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("delete employee: invalid id for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := getStore().Delete(id); err != nil {
		log.Printf("delete employee: failed for %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Employee Deleted")
}
