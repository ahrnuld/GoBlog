package model

import (
	"html/template"
	"time"
)

type Post struct {
	Id          int
	Title       string
	PostedAt    time.Time
	Content     string
	ContentHtml template.HTML
}
