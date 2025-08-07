package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := "root:Paruyr-2004-03-24@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Could not connect to database" + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	createTable()

	fmt.Println("Successfully connected to MySQL!")
}

func createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	userName VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
	);
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table." + err.Error())
	}

	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		data LONGBLOB,
		tags TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		author_id INTEGER,
		FOREIGN KEY(author_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		panic("Could not create posts table." + err.Error())
	}

	createPostMediaTable := "CREATE TABLE IF NOT EXISTS post_media (" +
		"id INT AUTO_INCREMENT PRIMARY KEY," +
		"post_id INT," +
		"media_url TEXT NOT NULL," +
		"media_type VARCHAR(20)," +
		"`order` INT DEFAULT 0," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE" +
		");"

	_, err = DB.Exec(createPostMediaTable)
	if err != nil {
		panic("Could not create post_medias table." + err.Error())
	}
}
