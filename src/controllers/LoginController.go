package controllers

import (
	"net/http"

	"github.com/tonychinwe/libraryone/src/authentication"
	"github.com/tonychinwe/libraryone/src/models"

	"encoding/json"

	"github.com/tonychinwe/libraryone/src/repository"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credential authentication.Credential
	json.NewDecoder(r.Body).Decode(&credential)
	var author models.Author
	err := repository.DB.First(&author, "Email =?", credential.Username).Error

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad credentials"))
	}

	if author.Password != credential.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad credentials. Password not correct"))
	}

	var token, err2 = authentication.GenerateJwt(credential)
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad credentials"))
	}
	json.NewEncoder(w).Encode(token)

}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credential authentication.Credential
	json.NewDecoder(r.Body).Decode(&credential)
	var admin models.Admin
	err := repository.DB.First(&admin, "Email =?", credential.Username).Error

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error() + "here 1"))
		return
	}

	if admin.Password != credential.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad credentials. Password not correct"))
		return
	}

	var token, err2 = authentication.GenerateJwt(credential)
	if err2 != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err2.Error() + token))
		return
	}
	var strToken = token + " super99"
	json.NewEncoder(w).Encode(strToken)

}
