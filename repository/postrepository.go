package repository

import (
	"GoBlog/model"
	"log"
)

func GetAllPosts() []model.Post {
	db := openConnection()

	// run query
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() // runs at end of function

	// process result
	var posts []model.Post

	for rows.Next() {
		var post model.Post
		rows.Scan(&post.Id, &post.Title, &post.PostedAt, &post.Content)
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return posts
}

func GetSinglePost(id int) model.Post {
	db := openConnection()

	rows, err := db.Query("SELECT * FROM post WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close() // runs at end of function

	// process result
	var post model.Post

	for rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.PostedAt, &post.Content)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return post
}

func CreatePost(post model.Post) {
	db := openConnection()

	_, err := db.Exec(`INSERT INTO post (title, content, posted_at) VALUES (?, ?, NOW())`, post.Title, post.Content)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePost(post model.Post) {
	db := openConnection()

	_, err := db.Exec(`UPDATE post SET title = ?, content = ? WHERE id = ?`, post.Title, post.Content, post.Id)
	if err != nil {
		log.Fatal(err)
	}
}
