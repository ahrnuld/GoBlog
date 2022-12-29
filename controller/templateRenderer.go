package controller

import (
	"html/template"
	"net/http"
)

func renderAdminTemplate(w http.ResponseWriter, templatefile string, data any) {
	files := []string{
		"./view/admin/base.html",
		templatefile,
	}

	ts, _ := template.ParseFiles(files...)

	ts.ExecuteTemplate(w, "base", data)
}
