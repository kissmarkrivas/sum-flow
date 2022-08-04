package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissmarkrivas/api-sum/db"
	"github.com/kissmarkrivas/api-sum/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request)  {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	
}

func GetUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user,params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user no fount"))
		return
	}
	db.DB.Model(&user).Association("Sum").Find(&user.Sum)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))	
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user,params["id"])
	if user.ID == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user no fount"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
