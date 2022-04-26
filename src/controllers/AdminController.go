package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/tonychinwe/libraryone/src/models"
	"github.com/tonychinwe/libraryone/src/repository"
)

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var admin models.Admin
	json.NewDecoder(r.Body).Decode(&admin)
	if (strings.Index(admin.Name, "#")) < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not permitted here"))
		return
	}
	adm := strings.Split(admin.Name, "#")
	if adm[1] != "super99" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You are not permitted here"))
		return
	}
	err := repository.DB.Create(&admin).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
