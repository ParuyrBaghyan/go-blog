package models

import (

	"go-blog/db"
	"time"
)

type Post struct {
	Id       int64
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	Tags     string    `json:"tags"`
	DateTime time.Time `json:"dateTime" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	AuthorId int       `json:"authorId"`
}

func GetPaginatedPosts(limit int, offset int) ([]Post, error) {
	query := "SELECT * FROM posts ORDER BY dateTime DESC LIMIT ? OFFSET ?"

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Tags, &post.DateTime, &post.AuthorId)
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
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Tags, &post.DateTime, &post.AuthorId)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Post) Save() error {
	query := "INSERT INTO posts(title,content,tags,dateTime,author_id) VALUES(?,?,?,?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(p.Title, p.Content, p.Tags, p.DateTime, p.AuthorId)
	if err != nil {
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.Id = postID

	return err
}

func (p *Post) Update() error {
	query := `
	UPDATE posts
	SET title = ?, content = ?, tags = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.Content, p.Tags, p.DateTime, p.Id)

	return err

}

func (p *Post) Delete() error {
	query := "DELETE FROM posts WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Id)

	return err
}
