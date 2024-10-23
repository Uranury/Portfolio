package handlers

import (
	"net/http"
)

func GetDepartments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func GetEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
