package handlers

import (
	"REST-API/db"
	"REST-API/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []models.Employee

	if err := db.DB.Preload("Department").Find(&employees).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee models.Employee
	EmployeeSSN, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid employee SSN", http.StatusBadRequest)
		return
	}
	if err := db.DB.Preload("Department").First(&employee, "ssn = ?", EmployeeSSN).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "employee not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// Save employee in the databse
	if err := db.DB.Create(&employee).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load the associated Department record
	if err := db.DB.Preload("Department").First(&employee, employee.SSN).Error; err != nil {
		http.Error(w, "failed to load department", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee models.Employee
	EmployeeSSN, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid employee SSN", http.StatusBadRequest)
		return
	}

	if err := db.DB.First(&employee, "ssn = ?", EmployeeSSN).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "employee not found", http.StatusNotFound)
		} else {
			log.Printf("Error checking employee: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := db.DB.Delete(&employee, "ssn = ?", EmployeeSSN).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
