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

func CreateSubject(w http.ResponseWriter, r *http.Request) {

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
	var subject models.Subject
	json.NewDecoder(r.Body).Decode(&subject)
	err := repository.DB.Create(&subject).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(subject)

}

func GetAllSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subjects []models.Subject
	repository.DB.Find(&subjects)
	json.NewEncoder(w).Encode(subjects)

}

func GetSubjectWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var subject models.Subject
	err := repository.DB.First(&subject, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(subject)
	}

}

func UpdateSubjectWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subject models.Subject
	json.NewDecoder(r.Body).Decode(&subject)
	var id = subject.ID
	err := repository.DB.First(&subject, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(subject)
}

func DeleteSubjectWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var subject models.Subject
	err := repository.DB.Delete(&subject, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(subject)
}
