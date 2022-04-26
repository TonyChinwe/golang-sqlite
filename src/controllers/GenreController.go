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

func CreateGenre(w http.ResponseWriter, r *http.Request) {

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
	var genre models.Genre
	json.NewDecoder(r.Body).Decode(&genre)
	err := repository.DB.Create(&genre).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(genre)

}

func GetAllGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var genres []models.Genre
	repository.DB.Find(&genres)
	json.NewEncoder(w).Encode(genres)

}

func GetGenreWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var genre models.Genre
	err := repository.DB.First(&genre, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(genre)
	}

}

func UpdateGenreWithId(w http.ResponseWriter, r *http.Request) {
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
	var genre models.Genre
	json.NewDecoder(r.Body).Decode(&genre)
	var id = genre.ID
	err := repository.DB.First(&genre, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(genre)
}

func DeleteGenreWithId(w http.ResponseWriter, r *http.Request) {
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
	var genre models.Genre
	err := repository.DB.Delete(&genre, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(genre)
}
