package database

import (
	"gorm.io/gorm"
)

// BlogPost represents a blog post record in the database.
// It includes fields for the title, author, and content, and embeds gorm.Model
// which provides ID, CreatedAt, UpdatedAt, and DeletedAt fields.
type Post struct { //table
	gorm.Model
	Title   string
	Author  string
	Content string
}

// dbBlog wraps a GORM database connection for performing blog-related operations.
// It acts as the concrete implementation of BlogManager.
type dbBlog struct { //blog database
	db *gorm.DB
}

// BlogManager defines the set of operations that can be performed on blog posts.
// Any struct implementing this interface must provide methods to create, read,
// update, delete, and search blog posts.
type BlogManager interface { //interface
	CreatePost(title string, author string, content string) (Post, error)
	ReadAllPosts() ([]Post, error)
	UpdatePost(id uint, newtitle string, content string) error
	DeletePost(title string) error
	SearchPost(author string, title string) ([]Post, error)
}
