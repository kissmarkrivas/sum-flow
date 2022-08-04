package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kissmarkrivas/api-sum/db"
	"github.com/kissmarkrivas/api-sum/models"
)

func GetSumsHandler(w http.ResponseWriter, r *http.Request)  {
	var sums []models.Sum
	db.DB.Find(&sums)
	json.NewEncoder(w).Encode(sums)
}

func CreateSumHandler(w http.ResponseWriter, r *http.Request)  {
	var sum models.Sum
	json.NewDecoder(r.Body).Decode(&sum)
	createdSum := db.DB.Create(&sum)
	err := createdSum.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&sum)
}

func GetSumHandler(w http.ResponseWriter, r *http.Request)  {
	var sum models.Sum
	params := mux.Vars(r)
	db.DB.First(&sum,params["id"])
	if sum.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sum Not Found"))
		return
	}
	json.NewEncoder(w).Encode(&sum)
}

func DeleteSumsHandler(w http.ResponseWriter, r *http.Request)  {
	var sum models.Sum
	params := mux.Vars(r)
	db.DB.First(&sum,params["id"])
	if sum.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Sum Not Found"))
		return
	}
	db.DB.Unscoped().Delete(&sum)
	w.WriteHeader(http.StatusNoContent)
}