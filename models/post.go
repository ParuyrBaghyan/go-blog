package models

import (
	"database/sql"
	"go-blog/db"
	"time"
)

type Post struct {
	Id       int64
	Title    string `binding:"required"`
	Content  string `binding:"required"`
	Data     sql.NullString
	Tags     string
	DateTime time.Time `binding:"required"`
	AuthorId int64
}

func GetAllPosts() ([]Post, error) {
	query := "SELECT * FROM posts"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Data, &post.Tags, &post.DateTime, &post.AuthorId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostById(postId int64) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"

	row := db.DB.QueryRow(query, postId)
	var post Post
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Data, &post.Tags, &post.DateTime, &post.AuthorId)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Post) Save() error {
	query := "INSERT INTO posts(title,content,data,tags,dateTime,author_id) VALUES(?,?,?,?,?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.Content, p.Data, p.Tags, p.DateTime, p.AuthorId)
	if err != nil {
		return err
	}

	return err
}
