package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonychinwe/libraryone/src/controllers"
)

func InitRouter() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/test", controllers.MyHandler).Methods("GET")

	//author
	r.HandleFunc("/create-author", controllers.CreateAuthor).Methods("POST")
	r.HandleFunc("/get-all-authors", controllers.GetAllAuthors).Methods("GET")
	r.HandleFunc("/get-author", controllers.GetAuthorWithId).Methods("GET")
	r.HandleFunc("/update-author", controllers.UpdateAuthorWithId).Methods("PUT")
	r.HandleFunc("/delete-author", controllers.DeleteAuthorWithId).Methods("DELETE")

	//book
	r.HandleFunc("/create-book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/get-all-book-author", controllers.GetAllBookByAuthor).Methods("GET")
	r.HandleFunc("/get-book-subject", controllers.GetAllBookBySubject).Methods("GET")
	r.HandleFunc("/get-book-category", controllers.GetAllBookByCategory).Methods("GET")
	r.HandleFunc("/get-all-books", controllers.GetAllBook).Methods("GET")
	r.HandleFunc("/get-book", controllers.GetBookWithId).Methods("GET")
	r.HandleFunc("/update-book", controllers.UpdateBookWithId).Methods("PUT")
	r.HandleFunc("/delete-book", controllers.DeleteBookWithId).Methods("DELETE")

	//category
	r.HandleFunc("/create-category", controllers.CreateCat).Methods("POST")
	r.HandleFunc("/get-all-category", controllers.GetAllCat).Methods("GET")
	r.HandleFunc("/get-category", controllers.GetCatWithId).Methods("GET")
	r.HandleFunc("/update-category", controllers.UpdateCatWithId).Methods("PUT")
	r.HandleFunc("/delete-category", controllers.DeleteCatWithId).Methods("DELETE")

	//genre
	r.HandleFunc("/create-genre", controllers.CreateGenre).Methods("POST")
	r.HandleFunc("/get-all-genre", controllers.GetAllGenre).Methods("GET")
	r.HandleFunc("/get-genre", controllers.GetGenreWithId).Methods("GET")
	r.HandleFunc("/update-genre", controllers.UpdateGenreWithId).Methods("PUT")
	r.HandleFunc("/delete-genre", controllers.DeleteGenreWithId).Methods("DELETE")

	//level
	r.HandleFunc("/create-level", controllers.CreateLevel).Methods("POST")
	r.HandleFunc("/get-all-level", controllers.GetAllLevel).Methods("GET")
	r.HandleFunc("/get-level", controllers.GetLevelWithId).Methods("GET")
	r.HandleFunc("/update-level", controllers.UpdateLevelWithId).Methods("PUT")
	r.HandleFunc("/delete-level", controllers.DeleteLevelWithId).Methods("DELETE")

	//subject
	r.HandleFunc("/create-subject", controllers.CreateSubject).Methods("POST")
	r.HandleFunc("/get-all-subject", controllers.GetAllSubject).Methods("GET")
	r.HandleFunc("/get-subject", controllers.GetSubjectWithId).Methods("GET")
	r.HandleFunc("/update-subject", controllers.UpdateSubjectWithId).Methods("PUT")
	r.HandleFunc("/delete-subject", controllers.DeleteSubjectWithId).Methods("DELETE")

	//login
	r.HandleFunc("/login-user", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/login-admin", controllers.LoginAdmin).Methods("POST")

	r.HandleFunc("/create-admin", controllers.CreateAdmin).Methods("POST")

	http.ListenAndServe(":8080", r)
}
