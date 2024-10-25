package handlers

import (
	"REST-API/db"
	"REST-API/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var departments []models.Department

	if err := db.DB.Find(&departments).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(departments)
}

func GetDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var department models.Department
	departmentCode, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid department code", http.StatusBadRequest)
		return
	}
	if err := db.DB.First(&department, "code = ?", departmentCode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "department not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(department)
}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var department models.Department

	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	if err := db.DB.Create(&department).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(department)
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var department models.Department
	departmentCode, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid department code", http.StatusBadRequest)
		return
	}

	if err := db.DB.First(&department, "code = ?", departmentCode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "department not found", http.StatusNotFound)
		} else {
			log.Printf("Error checking department: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := db.DB.Delete(&department, "code = ?", departmentCode).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extracting department ID from the URL parameters
	params := mux.Vars(r)
	departmentID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "invalid department code", http.StatusBadRequest)
		return
	}

	// Check if the department exists
	var department models.Department
	if err := db.DB.First(&department, "code = ?", departmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "department not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Query for employees in the specified department
	var employees []struct {
		SSN      int    `json:"ssn"`
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	}

	if err := db.DB.Model(&models.Employee{}).
		Joins("JOIN departments ON employees.department_id = departments.code").
		Where("departments.code = ?", departmentID).
		Select("employees.ssn, employees.name, employees.last_name").
		Find(&employees).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if any employees were found
	if len(employees) == 0 {
		http.Error(w, "no employees found for the specified department", http.StatusNotFound)
		return
	}

	// Encode the employee data as JSON and send the response
	json.NewEncoder(w).Encode(employees)
}

func GetTotalBudget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var department models.Department
	var totalBudget int64
	if err := db.DB.Model(&department).Select("SUM(budget)").Scan(&totalBudget).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "{total budget: %d}", totalBudget)
}
