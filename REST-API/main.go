package main

import (
	"REST-API/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/departments", handlers.GetDepartments).Methods("GET")
	router.HandleFunc("/departments/{id}", handlers.GetDepartment).Methods("GET")
	router.HandleFunc("/departments", handlers.CreateDepartment).Methods("POST")
	router.HandleFunc("/departments/{id}", handlers.DeleteDepartment).Methods("DELETE")
	router.HandleFunc("/departments/{id}/employees", handlers.GetEmployeesByDepartment).Methods("GET")

	router.HandleFunc("/employees/{id}", handlers.GetEmployee).Methods("GET")
	router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", handlers.DeleteEmployee).Methods("DELETE")

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
