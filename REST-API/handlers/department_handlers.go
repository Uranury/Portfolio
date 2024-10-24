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

/*
func GetEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
*/

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
