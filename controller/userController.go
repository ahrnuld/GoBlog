package controller

import (
	"GoBlog/model"
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type UserController struct {
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/user-create.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	user := model.User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, user.Username, user.Password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	_ = result

	tmpl.Execute(w, struct{ Success bool }{true})
}

type UserPageData struct {
	PageTitle string
	Users     []model.User
}

func (uc *UserController) UserJson(writer http.ResponseWriter, request *http.Request) {

	db, _ := sql.Open("mysql", "root:password@(127.0.0.1:3306)/dbname?parseTime=true")

	rows, _ := db.Query("SELECT id, username, password, created_at FROM users")
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
		users = append(users, user)
	}

	json.NewEncoder(writer).Encode(users)
}
