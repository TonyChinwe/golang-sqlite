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

func CreateCat(w http.ResponseWriter, r *http.Request) {
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
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)
	err := repository.DB.Create(&cat).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(cat)

}

func GetAllCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cats []models.Category
	repository.DB.Find(&cats)
	json.NewEncoder(w).Encode(cats)

}

func GetCatWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var cat models.Category
	err := repository.DB.First(&cat, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(cat)
	}

}

func UpdateCatWithId(w http.ResponseWriter, r *http.Request) {

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
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)
	var id = cat.ID
	err := repository.DB.First(&cat, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(cat)
}

func DeleteCatWithId(w http.ResponseWriter, r *http.Request) {

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
	var cat models.Category
	err := repository.DB.Delete(&cat, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(cat)
}
