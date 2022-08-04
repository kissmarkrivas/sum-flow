package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissmarkrivas/api-sum/db"
	"github.com/kissmarkrivas/api-sum/models"
	"github.com/kissmarkrivas/api-sum/routes"
)

func main() {
	db.DBConection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Sum{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users",routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}",routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users",routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}",routes.DeleteUserHandler).Methods("DELETE")
// sum routes

	r.HandleFunc("/sums",routes.GetSumsHandler).Methods("GET")
	r.HandleFunc("/sums/{id}",routes.GetSumHandler).Methods("GET")
	r.HandleFunc("/sums",routes.CreateSumHandler).Methods("POST")
	r.HandleFunc("/sums/{id}",routes.DeleteSumsHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
	
}
