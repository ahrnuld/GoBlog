package main

import (
	"GoBlog/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")

}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

// Middleware in go is easy, just wrap your function around the handler function
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "Hello, you've requested: %s of book %s\n", title, page)
}

func main() {

	// http only cookies for security
	store.Options.HttpOnly = true

	// router setup
	r := mux.NewRouter()

	var blogController controller.BlogController
	r.HandleFunc("/", blogController.Index)
	r.HandleFunc("/posts/create", blogController.Create)
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logging(logout))

	var userController controller.UserController
	r.HandleFunc("/users/create", userController.CreateUser)
	r.HandleFunc("/usersJSON", userController.UserJson)
	r.HandleFunc("/books/{title}/page/{page}", ReadBook).Methods("GET").Schemes("http")

	var adminController controller.AdminController
	r.HandleFunc("/admin", adminController.Index)
	r.HandleFunc("/admin/create", adminController.Create).Methods("GET")
	r.HandleFunc("/admin/createPost", adminController.Create).Methods("POST")
	r.HandleFunc("/admin/edit/{id}", adminController.Edit).Methods("GET")
	r.HandleFunc("/admin/edit/{id}", adminController.EditPost).Methods("POST")

	// TinyMCE & static files support
	r.PathPrefix("/tinymce").Handler(http.FileServer(http.Dir("./node_modules/")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":80", r)
}
