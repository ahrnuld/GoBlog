package controller

import (
	"GoBlog/model"
	"GoBlog/repository"
	"net/http"
)

type PageData struct {
	PageTitle string
	Posts     []model.Post
}

type BlogController struct {
}

func (c *BlogController) Index(w http.ResponseWriter, r *http.Request) {
	posts := repository.GetAllPosts()
	renderTemplate(w, "./view/blog.html", posts)
}
