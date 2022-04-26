package controllers

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	// "github.com/stretchr/testify/assert"
	// "github.com/tonychinwe/libraryone/src/models"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create-admin", CreateAdmin).Methods("POST")
	return router

}

func TestMyHandler(t *testing.T) {

}

func TestCreateAdmin(t *testing.T) {

	req, _ := http.NewRequest("GET", "/test", nil)
	handler := http.HandlerFunc(MyHandler)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	//Router().ServeHTTP(response, req)
	status := response.Code

	if status != http.StatusOK {

		t.Error("wrong status code")
	}

	//var data = []byte(`{"ID":1,"name":"name#super99","mail":"mail 1","password":"password 1"}`)
	// req, _ := http.NewRequest("POST", "/create-admin", bytes.NewBuffer(data))
	// handler := http.HandlerFunc(CreateAdmin)
	// response := httptest.NewRecorder()

	// Router().ServeHTTP(response, req)
	// status := response.Code

	// if status != http.StatusOK {

	// 	t.Error("wrong status code")
	// }

	// var admin models.Admin
	// fmt.Println(response.Body)
	// json.NewDecoder(io.Reader(response.Body)).Decode(&admin)
	// assert.NotNil(t, admin.ID)
	// assert.Equal(t, "name 1", admin.Name)
	// assert.Equal(t, "mail 1", admin.Email)

}
