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

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
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
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	err := repository.DB.Create(&author).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(author)

}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authors []models.Author
	repository.DB.Find(&authors)
	json.NewEncoder(w).Encode(authors)

}

func GetAuthorWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var author models.Author
	err := repository.DB.First(&author, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(author)
	}

}

func UpdateAuthorWithId(w http.ResponseWriter, r *http.Request) {

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
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	var id = author.ID
	err := repository.DB.First(&author, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthorWithId(w http.ResponseWriter, r *http.Request) {

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
	var author models.Author
	err := repository.DB.Delete(&author, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(author)
}
