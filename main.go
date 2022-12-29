package main

import (
	"GoBlog/controller"
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

func login(w http.ResponseWriter, r *http.Request) {
	// TODO: add login form and check in separate controller
	session, _ := store.Get(r, "auth-cookie")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth-cookie")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

// Middleware in go is easy, just wrap your function around the handler function
func checkAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "auth-cookie")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		f(w, r)
	}
}

func main() {

	// http only cookies for security
	store.Options.HttpOnly = true

	// router setup
	r := mux.NewRouter()

	var blogController controller.BlogController
	r.HandleFunc("/", blogController.Index)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	var userController controller.UserController
	r.HandleFunc("/users/create", userController.CreateUser)
	r.HandleFunc("/usersJSON", userController.UserJson)

	var adminController controller.AdminController
	r.HandleFunc("/admin", checkAuth(adminController.Index))
	r.HandleFunc("/admin/create", checkAuth(adminController.Create)).Methods("GET")
	r.HandleFunc("/admin/createPost", checkAuth(adminController.Create)).Methods("POST")
	r.HandleFunc("/admin/edit/{id}", checkAuth(adminController.Edit)).Methods("GET")
	r.HandleFunc("/admin/edit/{id}", checkAuth(adminController.EditPost)).Methods("POST")

	// TinyMCE & static files support
	r.PathPrefix("/tinymce").Handler(http.FileServer(http.Dir("./node_modules/")))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":80", r)
}
