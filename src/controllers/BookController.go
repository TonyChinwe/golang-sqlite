package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonychinwe/libraryone/src/authentication"
	"github.com/tonychinwe/libraryone/src/models"
	"github.com/tonychinwe/libraryone/src/repository"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if !authentication.IsAuthorized(r.Header.Get("Authorization")) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Bad Credentials"))
		return
	}
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	var subjectId = book.SubjectID
	var subject models.Subject
	err := repository.DB.First(&subject, "ID =?", subjectId).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		book.Subject = subject
	}

	var categoryId = book.CategoryID
	var category models.Category
	err = repository.DB.First(&category, "ID =?", categoryId).Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		book.Category = category
	}
	var genreId = book.GenreID
	var genre models.Genre
	err = repository.DB.First(&genre, "ID =?", genreId).Error
	if err == nil {
		book.Genre = genre
	}
	var levelId = book.Leveid
	var level models.Level
	err = repository.DB.First(&level, "Name =?", levelId).Error
	if err == nil {
		book.Level = level
	}

	repository.DB.Create(&book)
	json.NewEncoder(w).Encode(book)

}

func GetAllBookByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id = uint(r.FormValue("id")[0])

	var author models.Author
	err := repository.DB.First(&author, "ID =?", id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	var books []models.Book = author.Books
	json.NewEncoder(w).Encode(books)

}

func GetAllBookBySubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id = uint(r.FormValue("id")[0])

	var subject models.Subject
	err := repository.DB.First(&subject, "ID =?", id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	var books []models.Book
	repository.DB.Find(&books, "Subject =?", subject)
	json.NewEncoder(w).Encode(books)
}

func GetAllBookByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id = uint(r.FormValue("id")[0])

	var cat models.Category
	err := repository.DB.First(&cat, "ID =?", id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	var books []models.Book
	repository.DB.Find(&books, "Category =?", cat)
	json.NewEncoder(w).Encode(books)

}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	repository.DB.Find(&books)
	json.NewEncoder(w).Encode(books)

}

func GetBookWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var book models.Book
	err := repository.DB.First(&book, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else {

		json.NewEncoder(w).Encode(book)
	}

}

func UpdateBookWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	var id = book.ID
	err := repository.DB.First(&book, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(book)
}

func DeleteBookWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = mux.Vars(r)
	var book models.Book
	err := repository.DB.Delete(&book, path["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(book)
}
