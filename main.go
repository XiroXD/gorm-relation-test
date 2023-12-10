package main

import (
	"encoding/json"
	"fmt"

	"github.com/XiroXD/gorm-test/database"
	"github.com/XiroXD/gorm-test/models"
)

func init() {
	db := database.Connect()
	database.Migrate(db)
}

func main() {

	createPost := models.Post{
		Title:   "Hi mom!",
		Content: "This is my first post.",
		Author: models.Author{
			Name:  "John Doe",
			Email: "jY9kK@example.com",
		},
	}

	database.DBConn.Create(&createPost)

	
	var post models.Post
	var author models.Author

	database.DBConn.Preload("Author").First(&post, 1)
	database.DBConn.First(&author, post.Author)

	bytes, _ := json.Marshal(post)

	fmt.Println(string(bytes))
}
