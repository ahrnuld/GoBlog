package controller

import (
	"GoBlog/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminController struct {
}

func (c *AdminController) Index(w http.ResponseWriter, r *http.Request) {
	// get data
	posts := getAllPosts()
	// render view
	renderAdminTemplate(w, "./view/admin/list.html", posts)
}

func (c *AdminController) Create(w http.ResponseWriter, r *http.Request) {
	renderAdminTemplate(w, "./view/admin/create.html", nil)
}

// POST variant
func (c *AdminController) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := model.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}
	createPost(post)
	// forward to list view
	c.Index(w, r)
}

func (c *AdminController) Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	post := getSinglePost(id)
	renderAdminTemplate(w, "./view/admin/edit.html", post)
}

func (c *AdminController) EditPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	post := model.Post{
		Id:      id,
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}
	updatePost(post)
	// forward to list view
	c.Index(w, r)
}
