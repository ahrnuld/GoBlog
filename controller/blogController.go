package controller

import (
	"GoBlog/model"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	PageTitle string
	Posts     []model.Post
}

type BlogController struct {
}

func (c *BlogController) Index(w http.ResponseWriter, r *http.Request) {

	// open DB connection
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// run query
	rows, err := db.Query("SELECT * FROM post ORDER BY posted_at DESC")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() // runs at end of function

	// process result
	var posts []model.Post

	for rows.Next() {
		var post model.Post
		rows.Scan(&post.Id, &post.Title, &post.PostedAt, &post.Content)
		post.ContentHtml = template.HTML(post.Content)
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// create viewmodel
	data := PageData{
		PageTitle: "Random Title",
		Posts:     posts,
	}

	// parse view
	files := []string{
		"./view/base.html",
		"./view/blog.html",
	}

	ts, _ := template.ParseFiles(files...)

	ts.ExecuteTemplate(w, "base", data)
}

func (c *BlogController) Create(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("view/post-create.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	post := model.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	result, err := db.Exec(`INSERT INTO users (title, content, created_at) VALUES (?, ?, NOW())`, post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}

	_ = result

	tmpl.Execute(w, struct{ Success bool }{true})
}
