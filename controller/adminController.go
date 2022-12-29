package controller

import (
	"GoBlog/model"
	"GoBlog/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminController struct {
}

func (c *AdminController) Index(w http.ResponseWriter, r *http.Request) {
	// get data
	posts := repository.GetAllPosts()
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
	repository.CreatePost(post)
	// forward to list view
	c.Index(w, r)
}

func (c *AdminController) Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	post := repository.GetSinglePost(id)
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
	repository.UpdatePost(post)
	// forward to list view
	c.Index(w, r)
}
