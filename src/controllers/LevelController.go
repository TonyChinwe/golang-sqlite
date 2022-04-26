package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/tonychinwe/libraryone/src/authentication"
	"github.com/tonychinwe/libraryone/src/models"
	"github.com/tonychinwe/libraryone/src/repository"
)

func CreateLevel(w http.ResponseWriter, r *http.Request) {
	var token1 = r.Header.Get("Authorization")
	t1 := strings.Split(token1, " ")
	if len(t1) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}

	if t1[1] != "super99" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials"))
		return
	}

	if !authentication.IsAuthorized(t1[0]) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var level models.Level
	json.NewDecoder(r.Body).Decode(&level)
	err := repository.DB.Create(&level).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(level)

}

func GetAllLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var levels []models.Level
	repository.DB.Find(&levels)
	json.NewEncoder(w).Encode(levels)

}

func GetLevelWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var level models.Level
	err := repository.DB.First(&level, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(level)
	}

}

func UpdateLevelWithId(w http.ResponseWriter, r *http.Request) {
	var token1 = r.Header.Get("Authorization")
	t1 := strings.Split(token1, " ")
	if len(t1) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}

	if t1[1] != "super99" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials"))
		return
	}

	if !authentication.IsAuthorized(t1[0]) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var level models.Level
	json.NewDecoder(r.Body).Decode(&level)
	var id = level.ID
	err := repository.DB.First(&level, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(level)
}

func DeleteLevelWithId(w http.ResponseWriter, r *http.Request) {
	var token1 = r.Header.Get("Authorization")
	t1 := strings.Split(token1, " ")
	if len(t1) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}

	if t1[1] != "super99" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials"))
		return
	}

	if !authentication.IsAuthorized(t1[0]) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials "))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var level models.Level
	err := repository.DB.Delete(&level, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(level)
}
